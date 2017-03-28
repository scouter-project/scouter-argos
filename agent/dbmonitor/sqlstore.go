package dbmonitor

GlobalStatusSQL := "SELECT variable_name, variable_value " + "FROM information_schema.global_status "
			+ "WHERE variable_name IN " + "('THREADS_RUNNING','THREADS_CONNECTED','INNODB_BUFFER_POOL_PAGES_TOTAL',"
			+ "'INNODB_BUFFER_POOL_PAGES_DATA','INNODB_BUFFER_POOL_PAGES_FREE','INNODB_BUFFER_POOL_PAGES_DIRTY',"
			+ "'INNODB_ROW_LOCK_CURRENT_WAITS','COM_SELECT','QCACHE_HITS','COM_INSERT',"
			+ "'COM_UPDATE','COM_DELETE','ABORTED_CLIENTS','ABORTED_CONNECTS',"
			+ "'CONNECTIONS','COM_CALL_PROCEDURE','INNODB_BUFFER_POOL_READS',"
			+ "'KEY_READS','KEY_READ_REQUESTS','QCACHE_INSERTS','THREADS_CREATED',"
			+ "'QUESTIONS','COM_COMMIT','COM_ROLLBACK','SLOW_QUERIES','SELECT_FULL_JOIN',"
			+ "'OPENED_TABLE_DEFINITIONS','CREATED_TMP_DISK_TABLES','CREATED_TMP_TABLES',"
			+ "'CREATED_TMP_FILES','INNODB_ROW_LOCK_WAITS','TABLE_LOCKS_IMMEDIATE',"
			+ "'TABLE_LOCKS_WAITED','THREADS_CACHED','SELECT_FULL_RANGE_JOIN',"
			+ "'SELECT_RANGE','SELECT_RANGE_CHECK','SELECT_SCAN','SORT_MERGE_PASSES',"
			+ "'SORT_PRIORITY_QUEUE_SORTS','SORT_RANGE','SORT_ROWS','SORT_SCAN',"
			+ "'INNODB_PAGES_CREATED','INNODB_PAGES_READ','INNODB_PAGES_WRITTEN',"
			+ "'INNODB_BUFFER_POOL_READ_REQUESTS','INNODB_ROWS_DELETED','INNODB_ROWS_INSERTED',"
			+ "'INNODB_ROWS_READ','INNODB_ROWS_UPDATED','BYTES_SENT','BYTES_RECEIVED',"
			+ "'KEY_WRITE_REQUESTS','KEY_WRITES','KEY_BLOCKS_USED','KEY_BLOCKS_NOT_FLUSHED',"
			+ "'INNODB_MUTEX_OS_WAITS','INNODB_MUTEX_SPIN_ROUNDS','INNODB_MUTEX_SPIN_WAITS',"
			+ "'INNODB_LOG_WAITS','INNODB_LOG_WRITE_REQUESTS','INNODB_LOG_WRITES',"
			+ "'INNODB_OS_LOG_WRITTEN','INNODB_ROW_LOCK_TIME','INNODB_DATA_FSYNCS','INNODB_DATA_READS','INNODB_DATA_WRITES',"
			+ "'INNODB_CHECKPOINT_AGE','QCACHE_NOT_CACHED','QCACHE_QUERIES_IN_CACHE','QCACHE_LOWMEM_PRUNES')"

InnodbLockSQL := "SELECT r.trx_wait_started AS wait_started,  " + " rl.lock_table AS locked_table,      "
			+ " rl.lock_index AS locked_index,      " + " rl.lock_type AS locked_type,      "
			+ " r.trx_started AS waiting_trx_started,      " + " r.trx_mysql_thread_id AS waiting_pid,      "
			+ " IFNULL(r.trx_query,'n/a') AS waiting_query,      " + " rl.lock_mode AS waiting_lock_mode,      "
			+ " b.trx_mysql_thread_id AS blocking_pid,      " + " IFNULL(b.trx_query,'n/a') AS blocking_query,      "
			+ " bl.lock_mode AS blocking_lock_mode,      " + " b.trx_started AS blocking_trx_started,      "
			+ " 0 AS status   " + " FROM information_schema.innodb_lock_waits w      "
			+ " INNER JOIN information_schema.innodb_trx b    ON b.trx_id = w.blocking_trx_id      "
			+ " INNER JOIN information_schema.innodb_trx r    ON r.trx_id = w.requesting_trx_id      "
			+ " INNER JOIN information_schema.innodb_locks bl ON bl.lock_id = w.blocking_lock_id      "
			+ " INNER JOIN information_schema.innodb_locks rl ON rl.lock_id = w.requested_lock_id     "
			+ " UNION ALL   " + " SELECT CURRENT_TIMESTAMP AS wait_started,  "
			+ " root.lock_table AS locked_table,      " + " root.lock_index AS locked_index,      "
			+ " root.lock_type AS locked_type,      " + " CURRENT_TIMESTAMP AS waiting_trx_started,      "
			+ " 0 AS waiting_pid,      " + " '' AS waiting_query,      " + " '' AS waiting_lock_mode,      "
			+ " r.trx_mysql_thread_id AS blocking_pid,      " + " IFNULL(r.trx_query,'n/a') AS blocking_query,      "
			+ " root.lock_mode AS blocking_lock_mode,      " + " r.trx_started AS blocking_trx_started,      "
			+ " 1 AS status  " + " FROM information_schema.innodb_trx r   "
			+ " INNER JOIN  (SELECT t1.lock_trx_id,t1.lock_mode,t1.lock_type,t1.lock_table,t1.lock_index  "
			+ "	      FROM information_schema.innodb_locks t1   "
			+ "	      LEFT JOIN information_schema.innodb_lock_waits t2 ON t1.lock_trx_id = t2.requesting_trx_id  "
			+ "	      WHERE t2.requesting_trx_id IS NULL ) root  " + "ON r.trx_id = root.lock_trx_id"

InnodbLockRootCountSQL := "SELECT count(*) FROM information_schema.innodb_locks t1 "
			+ "LEFT JOIN information_schema.innodb_lock_waits t2 " + "ON t1.lock_trx_id = t2.requesting_trx_id "
			+ "WHERE t2.requesting_trx_id IS NULL"

PROCESS_LIST_USE_V5 := "SELECT id, user, host, db, command, state, time, info , "
			+ "-1 as memory_used, -1 as examined_rows " + "FROM information_schema.processlist "
			+ "WHERE state not in ('Sleep', '') " + "AND user != 'system user' " + "ORDER BY db, time"

PROCESS_LIST_USE_v10 := "SELECT id, user, host, db, command, state, time, info , "
			+ "memory_used, examined_rows " + "FROM information_schema.processlist "
			+ "WHERE state not in ('Sleep', '') " + "AND user != 'system user' " + "ORDER BY db, time"

PROCESS_LIST_USE_P_F := "SELECT processlist_id as id, processlist_user as user, "
			+ "processlist_host as host,  processlist_db as db, "
			+ "processlist_command as command, processlist_state as state, "
			+ "processlist_time as time, processlist_info as info ," + "-1 as memory_used, -1 as examined_rows "
			+ "FROM performance_schema.threads " + "WHERE processlist_state not in ('Sleep', '') "
			+ "AND processlist_user != 'system user' " + "ORDER BY processlist_db, processlist_time";

SLAVE__STATUS := "SHOW SLAVE STATUS";

SLOW_QUERY_MAX_TIME := "SELECT max(start_time) from mysql.slow_log";

SLOW_QUERY := "SELECT start_time,user_host,query_time,lock_time, "
							 + "rows_sent,rows_examined,db,sql_text,thread_id FROM mysql.slow_log "
			                 + "WHERE start_time > ? ORDER BY start_time";




