# apollo configuration file

# 降级
[degrade]
mysql = 0
redis = 1
tusd = 0

[degrade.fileSharing]
push_server = 2
key1 = 1

# 熔断
[circuit]
default_timeout = 1000 # DefaultTimeout is how long to wait for command to complete, in milliseconds
default_max_concurrent_requests = 10 # DefaultMaxConcurrent is how many commands of the same type can run at the same time
default_request_volume_threshold = 20 # DefaultVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health
default_sleep_window = 5000 # DefaultSleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery
default_error_percent_threshold = 50 # DefaultErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests

[circuit.outgo.push]
timeout = 0
max_concurrent_requests = 5
request_volume_threshold = 5
sleep_window = 2000
error_percent_threshold = 3

[circuit.flow.data_report]
timeout = 0
max_concurrent_requests = 5
request_volume_threshold = 5
sleep_window = 2000
error_percent_threshold = 3


# others
[section_1]
key1 = 1
key2 = 2

[section_2]
key1 = 1
key2 = 2