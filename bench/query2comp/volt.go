package query2comp

import (
	"strconv"
	"testing"

	"github.com/akmonengine/volt"
	"github.com/mlange-42/go-ecs-benchmarks/bench/comps"
)

type voltConfig = volt.ComponentConfig[volt.ComponentInterface]

func runVolt(b *testing.B, n int) {
	b.StopTimer()
	world := volt.CreateWorld(1024)

	volt.RegisterComponent[comps.Position](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})
	volt.RegisterComponent[comps.Velocity](world, &voltConfig{BuilderFn: func(component any, configuration any) {}})

	for i := 0; i < n; i++ {
		e := world.CreateEntity(strconv.Itoa(i))
		volt.AddComponents2(world, e, comps.Position{}, comps.Velocity{})
	}

	query := volt.CreateQuery2[comps.Position, comps.Velocity](world, volt.QueryConfiguration{})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for result := range query.Foreach(nil) {
			pos := result.A
			vel := result.B
			pos.X += vel.X
			pos.Y += vel.Y
		}
	}
}
