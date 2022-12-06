package _1029_Two_City_Scheduling_

import "testing"

func Test_twoCitySchedCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				costs: [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
			},
			want: 110,
		},
		{
			name: "test2",
			args: args{
				costs: [][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}},
			},
			want: 3086,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoCitySchedCost(tt.args.costs); got != tt.want {
				t.Errorf("twoCitySchedCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
