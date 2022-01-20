package main

import "sort"

type Pair struct {
	start int
	end   int
}

type Job struct {
	pair   Pair
	profit int
}

type Jobs []Job

func JobScheduling(startTimes []int, endTimes []int, profits []int) int {
	var jobs Jobs

	for i := 0; i < len(profits); i++ {
		pair := Pair{start: startTimes[i], end: endTimes[i]}
		job := Job{pair: pair, profit: profits[i]}
		jobs = append(jobs, job)
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].pair.start < jobs[j].pair.start
	})

	return findMaxProfitFromJobs(jobs)
}

func findMaxProfitFromJobs(jobs []Job) int {
	n := len(jobs)
	maxProfit := 0

	pq := make(PriorityQueue, n)

	for i := 0; i < n; i++ {
		start := jobs[i].pair.start
		end := jobs[i].pair.end
		profit := jobs[i].profit

		for pq.IsEmpty() && start >= pq.top().pair.start {
			maxProfit = max(maxProfit, pq.top().profit)
			pq.Pop()
		}

		pq.Push(Job{start: start, end: end, profit: profit + maxProfit})
	}

	for pq.IsEmpty() {
		maxProfit = max(maxProfit, pq.Top().profit)
		pq.Remove()
	}

	return maxProfit
}

func main() {

}

/*

We have n jobs, where every job is scheduled to be done from startTime[i] to endTime[i], obtaining a profit of profit[i].

You're given the startTime, endTime and profit arrays, return the maximum profit you can take such that there are no two jobs in the subset with overlapping time range.

If you choose a job that ends at time X you will be able to start another job that starts at time X.

Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
Output: 120
Explanation: The subset chosen is the first and fourth job.
Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.

Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
Output: 150
Explanation: The subset chosen is the first, fourth and fifth job.
Profit obtained 150 = 20 + 70 + 60.
*/
