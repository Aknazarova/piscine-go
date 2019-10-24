package piscine

func AppendRange(min, max int) []int {
        if min >= max {
                var a []int = nil
                return a
        }else{
                a := []int{}
		
                for i:= min;i < max;i++ {
                        a = append(a,i)
                }
                return a
        }
}
