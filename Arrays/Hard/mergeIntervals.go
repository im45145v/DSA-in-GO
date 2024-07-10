func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }

    sortIntervals(intervals)

    mergedIntervals := make([][]int, 0, len(intervals))
    mergedIntervals = append(mergedIntervals, intervals[0])

    for _, interval := range intervals[1:] {
        if top := mergedIntervals[len(mergedIntervals)-1]; interval[0] > top[1] {
            mergedIntervals = append(mergedIntervals, interval)
        } else if interval[1] > top[1] {
            top[1] = interval[1]
        }
    }

    return mergedIntervals
}

func sortIntervals(intervals [][]int) {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
}
