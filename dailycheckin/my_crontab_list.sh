#必须要的默认定时任务请勿删除
0 */12 * * * sh /dailycheckin/default_task.sh >> /dailycheckin/default_task.log 2>&1
# 每天的 23:50 分清理一次日志
50 23 */2 * * rm -rf /dailycheckin/logs/*.log


##############每日签到一次任务##############
# 每日签到(every 8 hours 执行一次)
45 */8 * * * cd /dailycheckin && dailycheckin >> /dailycheckin/dailycheckin.log 2>&1
