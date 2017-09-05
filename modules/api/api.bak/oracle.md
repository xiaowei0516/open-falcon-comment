###Oracle监控


**对Oracle的监控支持OCI和ODBC两种方式：OCI是Oracle Call Interface的简称，即Oracle调用接口，通过OCI方式监测需要下载安装Oracle的OCI程序；通过ODBC进行监测需要在监测服务器上安装Oracle的管理端。**

**Oracle监控指标列表**

|监测点	 | 监测指标 | falcon-metric |	指标含义 |
|------|:----------:|:--------------:|-------:|
| 数据库性能	  |游标数(个)	    | databaseperformance.cursor     |数据库当前游标总数（显式游标+隐式游标）
|数据库性能       |会话数(个)	    | databaseperformance.sessions    |当前连接到实例的Sessions数
|数据库性能       |数据库锁总数(个)	    | databaseperformance.lock        |数据库中各种锁的总数
|数据库性能       |死锁总数(个)	    | databaseperformance.deadlock    |数据库中死锁的总数 
|数据库性能       |缓冲池命中率(%)	    | databaseperformance.cacheHitPercent | 数据块在数据缓冲区中的命中率
|数据库性能       |库cache命中率(%)    | databaseperformance.libHitPercent | 所发送的SQL语句在library cache中能找到它的执行方案的机率
| 系统当前状态	|Db block gets()    | system.blockRead                  | 在当前读模式下所读的块数
|系统当前状态     |Consistent gets(K块) | system.consistentRead           |在一致读模式下所读的块数，包括从回滚段读的块数
|系统当前状态     |Physical reads()     | system.diskRead                 |从磁盘上读取数据块的数量
|系统当前状态     |Redo entries()	      | system.redoCopy                 |Redo条目复制到Redo日志缓冲区的次数
|系统当前状态     |Redo buff allocation retry()	| system.redoBuff       |统计由于没有可用的redo log buffer，而不得不使一个为了复制新的redo entries进入redo log buffer的用户进程进入等待状态的次数
|系统当前状态     |Cache命中率(%)	                | system.cache.HitPercent | 未发生物理文件读取的数据请求在所有数据请求中所占比例
| 逻辑读Top10	|逻辑读Top N语句()	        | logicread.topN      | 从内存读到的数据块的Top N语句
|逻辑读Top10     |逻辑读Top N次数(次)	        | logicread.topNnum   |从内存读到的数据块的Top N次数
| 物理读Top10	|物理读Top N语句()	        | physicalread.topN   |从磁盘上读取数据块的Top N语句
|物理读Top10     |物理读Top N次数(次)	        | physicalread.topNnum |从磁盘上读取数据块的Top N次数
| 执行Top10	|执行Top N语句()	                | exec.topN            |执行Top N语句
|执行Top10       |执行Top N次数(次)	        | exec.topNnum         |执行Top N次数
|Oracle DBTableSpace      |可用百分比(%)	        |dbspace.usagePercent   | 可用空间/总空间大小
|Oracle DBTableSpace      |可用空间大小(M)	|dbspace.usageM         |表可用空间
|Oracle DBTableSpace      |总空间大小(M)	        |dbspace.totalM         |表总空间大小
|Oracle DBTableSpace      |最大碎片数()	        |dbspace.maxChip       |表空间最大碎片数
|Oracle DBTableSpace      |碎片总数()	        |dbspace.chip         |表空间碎片总数
|Oracle DBTableSpace      |数据文件()	        |dbspace.filePath     |表空间数据文件路径
|Oracle DBTableSpace      |读磁盘次数(次)	        |dbspace.readDisk  |表空间读磁盘次数
|Oracle DBTableSpace      |写磁盘次数(次)	        |dbspace.writeDisk    |表空间写磁盘次数
|Oracle namespace         |活动代理程序数(个)	|db.ativeAgent        |正在监测的数据库管理实例中已注册代理程序的数据(协调代理程序和子代理程序)
|Oracle namespace         |空闲代理程序数(个)	|db.freeAgent         |代理池程序池中当前未分配给应用程序而“空闲”的代理程序数
|Oracle namespace         |最大并行协调代理程序数(个)|db.maxAgent        |工作一段时间的协调代理的最大数量
|Oracle namespace         |代理程序创建频率(%)	 |db.createAgentPercent | 由于代理池为空而创建的代理数目/从代理池分配的代理的数目
