package statsd

import (
	"../schema"

	"gopkg.in/alexcesaro/statsd.v2"
	"time"
)

func NewStatsd(address, prefix string) (*statsd.Client, error) {
	return statsd.New(statsd.Address(address), statsd.Prefix(prefix))
}

func Save(client *statsd.Client, statistics *schema.Statistics) error {
	sendInt(client, "transactions", statistics.Transactions)
	sendFloat(client, "availability", statistics.Availability)
	sendDuration(client, "elapsed_time", statistics.ElapsedTime)
	sendFloat(client, "data_transfered", statistics.TransferredData)
	sendDuration(client, "response_time", statistics.ResponseTime)
	sendFloat(client, "transaction_rate", statistics.TransactionRate)
	sendFloat(client, "throughput", statistics.Throughput)
	sendFloat(client, "concurrency", statistics.Concurrency)
	sendInt(client, "successfull_transactions", statistics.SuccessfulTransactions)
	sendInt(client, "failed_transactions", statistics.FailedTransactions)
	sendDuration(client, "longest_transaction", statistics.LongestTransaction)
	sendDuration(client, "shortest_transaction", statistics.ShortestTransaction)

	sendDuration(client, "response_time.p10", statistics.Percentiles.P10)
	sendDuration(client, "response_time.p20", statistics.Percentiles.P20)
	sendDuration(client, "response_time.p30", statistics.Percentiles.P30)
	sendDuration(client, "response_time.p40", statistics.Percentiles.P40)
	sendDuration(client, "response_time.p50", statistics.Percentiles.P50)
	sendDuration(client, "response_time.p60", statistics.Percentiles.P60)
	sendDuration(client, "response_time.p70", statistics.Percentiles.P70)
	sendDuration(client, "response_time.p80", statistics.Percentiles.P80)
	sendDuration(client, "response_time.p90", statistics.Percentiles.P90)

	return nil
}

func sendDuration(client *statsd.Client, bucket string, duration time.Duration) {
	client.Timing(bucket, int(duration / time.Millisecond))
}

func sendFloat(client *statsd.Client, bucket string, value float64) {
	client.Gauge(bucket, value)
}

func sendInt(client *statsd.Client, bucket string, value int) {
	client.Gauge(bucket, value)
}