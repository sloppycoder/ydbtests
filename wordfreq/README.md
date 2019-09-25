## Test Steps

*  ``` env | grep ydb | wc -l``` should return 0
*  ``` source /usr/local/lib/yottadb/r128/ydb_env_set```
*  ``` env | grep ydb | wc -l``` should be more than 0
*  ``` gcc $(pkg-config --libs --cflags yottadb) -o wordfreq wordfreq.c -lyottadb```
*  ``` ./wordfreq < wordfreq_input.txt | tee output.txt``` should not print any error
*  ``` source /usr/local/lib/yottadb/r128/ydb_env_set```
*  ``` env | grep ydb | wc -l``` should return 0

