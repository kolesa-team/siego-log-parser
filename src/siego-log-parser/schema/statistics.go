package schema

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

/*
<result>
	<transactions>10</transactions>
	<availability>100.00%</availability>
	<elapsed_time>1.5694s</elapsed_time>
	<data_transferred>0.3997Mb</data_transferred>
	<response_time>0.1569s</response_time>
	<transaction_rate>6.3720/s</transaction_rate>
	<throughput>267076.4485Mb/s</throughput>
	<concurrency>7.9790</concurrency>
	<successful_transactions>10</successful_transactions>
	<failed_transactions/>
	<longest_transaction>1.5674s</longest_transaction>
	<shortest_transaction>0.6036s</shortest_transaction>
	<response_codes>
		<http_200>10</http_200>
	</response_codes>
	<percentiles>
		<p10>0.6453s</p10>
		<p20>1.2755s</p20>
		<p30>1.3174s</p30>
		<p40>1.3195s</p40>
		<p50>1.3219s</p50>
		<p60>1.3479s</p60>
		<p70>1.5610s</p70>
		<p80>1.5625s</p80>
		<p90>1.5674s</p90>
	</percentiles>
</result>
*/

type rawPercentile struct {
	P10 string `xml:"p10,omitempty"`
	P20 string `xml:"p20,omitempty"`
	P30 string `xml:"p30,omitempty"`
	P40 string `xml:"p40,omitempty"`
	P50 string `xml:"p50,omitempty"`
	P60 string `xml:"p60,omitempty"`
	P70 string `xml:"p70,omitempty"`
	P80 string `xml:"p80,omitempty"`
	P90 string `xml:"p90,omitempty"`
}

type rawStatistics struct {
	XMLName                xml.Name      `xml:"result"`
	Transactions           string        `xml:"transactions,omitempty"`
	Availability           string        `xml:"availability,omitempty"`
	ElapsedTime            string        `xml:"elapsed_time,omitempty"`
	TransferredData        string        `xml:"data_transferred,omitempty"`
	ResponseTime           string        `xml:"response_time,omitempty"`
	TransactionRate        string        `xml:"transaction_rate,omitempty"`
	Throughput             string        `xml:"throughput,omitempty"`
	Concurrency            string        `xml:"concurrency,omitempty"`
	SuccessfulTransactions string        `xml:"successful_transactions,omitempty"`
	FailedTransactions     string        `xml:"failed_transactions,omitempty"`
	LongestTransaction     string        `xml:"longest_transaction,omitempty"`
	ShortestTransaction    string        `xml:"shortest_transaction,omitempty"`
	Percentiles            rawPercentile `xml:"percentiles,omitempty"`
}

// Percentile — represents percentiles for response time
type Percentile struct {
	P10 time.Duration
	P20 time.Duration
	P30 time.Duration
	P40 time.Duration
	P50 time.Duration
	P60 time.Duration
	P70 time.Duration
	P80 time.Duration
	P90 time.Duration
}

// Statistics — represents statistics structure
type Statistics struct {
	Transactions           int
	Availability           float64
	ElapsedTime            time.Duration
	TransferredData        float64
	ResponseTime           time.Duration
	TransactionRate        float64
	Throughput             float64
	Concurrency            float64
	SuccessfulTransactions int
	FailedTransactions     int
	LongestTransaction     time.Duration
	ShortestTransaction    time.Duration
	Percentiles            Percentile
}

// ParseStatistics — parses byte array to statistics structure
func ParseStatistics(data []byte) (*Statistics, error) {
	result := Statistics{}
	statistics := rawStatistics{}

	err := xml.Unmarshal(data, &statistics)
	if err != nil {
		return &result, err
	}

	statistics.Availability = strings.Replace(statistics.Availability, "%", "", -1)
	statistics.TransferredData = strings.Replace(statistics.TransferredData, "Mb", "", -1)
	statistics.TransactionRate = strings.Replace(statistics.TransactionRate, "/s", "", -1)
	statistics.Throughput = strings.Replace(statistics.Throughput, "Mb/s", "", -1)

	result.Transactions, _ = strconv.Atoi(statistics.Transactions)
	result.Availability, _ = strconv.ParseFloat(statistics.Availability, 64)
	result.ElapsedTime, err = time.ParseDuration(statistics.ElapsedTime)
	result.TransferredData, err = strconv.ParseFloat(statistics.TransferredData, 64)
	result.ResponseTime, err = time.ParseDuration(statistics.ResponseTime)
	result.TransactionRate, err = strconv.ParseFloat(statistics.TransactionRate, 64)
	result.Throughput, err = strconv.ParseFloat(statistics.Throughput, 64)
	result.Concurrency, err = strconv.ParseFloat(statistics.Concurrency, 64)

	if statistics.SuccessfulTransactions != "" {
		result.SuccessfulTransactions, err = strconv.Atoi(statistics.SuccessfulTransactions)
	}

	if statistics.FailedTransactions != "" {
		result.FailedTransactions, err = strconv.Atoi(statistics.FailedTransactions)
	}

	result.LongestTransaction, err = time.ParseDuration(statistics.LongestTransaction)
	result.ShortestTransaction, err = time.ParseDuration(statistics.ShortestTransaction)

	result.Percentiles.P10, err = time.ParseDuration(statistics.Percentiles.P10)
	result.Percentiles.P20, err = time.ParseDuration(statistics.Percentiles.P20)
	result.Percentiles.P30, err = time.ParseDuration(statistics.Percentiles.P30)
	result.Percentiles.P40, err = time.ParseDuration(statistics.Percentiles.P40)
	result.Percentiles.P50, err = time.ParseDuration(statistics.Percentiles.P50)
	result.Percentiles.P60, err = time.ParseDuration(statistics.Percentiles.P60)
	result.Percentiles.P70, err = time.ParseDuration(statistics.Percentiles.P70)
	result.Percentiles.P80, err = time.ParseDuration(statistics.Percentiles.P80)
	result.Percentiles.P90, err = time.ParseDuration(statistics.Percentiles.P90)

	return &result, err
}
