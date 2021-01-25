package dispatcher

import "testing"

const (
	signal_1 = Signal(iota)
)

type TestDis struct{}

func (t TestDis) Dis(signal Signal, args ...interface{}) {}

func TestAddDisWithSignal(t *testing.T) {
	type args struct {
		dis     IDis
		signals []Signal
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "pass",
			args: args{
				dis:     TestDis{},
				signals: []Signal{signal_1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddDisWithSignal(tt.args.dis, tt.args.signals...)
		})
	}
}

func TestDelDis(t *testing.T) {
	type args struct {
		dis IDis
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "pass",
			args: args{
				dis: TestDis{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelDis(tt.args.dis)
		})
	}
}

func TestDelDisWithSignal(t *testing.T) {
	type args struct {
		signals []Signal
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "pass",
			args: args{
				signals: []Signal{signal_1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DelDisWithSignal(tt.args.signals...)
		})
	}
}

func TestDispatcher(t *testing.T) {
	AddDisWithSignal(TestDis{}, signal_1)
	type args struct {
		signal Signal
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass",
			args: args{
				signal: signal_1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Dispatcher(tt.args.signal, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Dispatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
