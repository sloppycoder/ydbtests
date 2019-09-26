### Helloworld of Golang and YottaDB

```
source /usr/local/lib/yottadb/r128/ydb_env_set

go run main.go

source /usr/local/lib/yottadb/r128/ydb_env_unset

```

### Test with replication setup

```
../repl_procedures

# setup A site as source
source ./ydbenv A r128
./db_create
./repl_setup

# start replication process at source site
./originating_start A B 4001

# open another terminal window, get into same directory
# set up B site as receiver site

source ./ydbenv B r128
./db_create
./repl_setup

# start replication receiver
./replicating_start_with_filter B 4001
tail -f B/receiver.log

# write somet to A site database
# open another terminal window
go build -o ../repl_procedures/tester .

cd ../repl_procedure
source ./ydbenv A r128
./tester
 
...
check B/filter.log to see any events are received!

```