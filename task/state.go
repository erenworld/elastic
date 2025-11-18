package task

func Contains(states []State, state) bool {
    for _, s := range states {
        if s == state {
            return true
        }
    }
    return false
}