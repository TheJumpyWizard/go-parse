%{
package main
%}

%token SELECT INSERT INTO VALUES FROM WHERE AND OR NOT GROUP BY HAVING ORDER BY ASC DESC LIMIT OFFSET ID NUMBER STRING COMMA LPAREN RPAREN EQ NEQ LT GT LTE GTE

%left OR
%left AND
%left EQ NEQ LT GT LTE GTE

%start statement

%%

statement: select_statement     {$$.result = $1}
         | insert_statement     {$$.result = $1}
         ;

select_statement: SELECT column_list FROM table_name (where_clause)? (group_by_clause)? (having_clause)? (order_by_clause)? (limit_clause)? (offset_clause)?
                 {
                     tableName := $4.(string)
                     columns := $2.([]string)
                     columnsMap := make(map[string]int)
                     for i, col := range columns {
                         if _, ok := tableColumns[tableName][col]; !ok {
                             return nil, fmt.Errorf("column %s does not exist in table %s", col, tableName)
                         }
                         columnsMap[col] = i
                     }
                     $$.result = &selectStatement{
                         Columns:    columns,
                         TableName:  tableName,
                         Where:      $5,
                         GroupBy:    $6.([]string),
                         Having:     $7,
                         OrderBy:    $8.([]string),
                         Ordering:   $9,
                         Limit:      $10,
                         Offset:     $11,
                         ColumnsMap: columnsMap,
                     }
                 }
                 ;


insert_statement: INSERT INTO table_name LPAREN column_list RPAREN VALUES LPAREN value_list RPAREN
                 {
                     tableName := $3.(string)
                     columnNames := $5.([]string)
                     values := $8.([][]interface{})
                     numColumns := len(columnNames)
                     for _, row := range values {
                         if len(row) != numColumns {
                             return nil, fmt.Errorf("number of columns in row does not match number of columns in insert statement")
                         }
                     }
                     $$.result = &insertStatement{
                         TableName:   tableName,
                         ColumnNames: columnNames,
                         Values:      values,
                         NumColumns:  numColumns,
                     }
                 }
                 ;
