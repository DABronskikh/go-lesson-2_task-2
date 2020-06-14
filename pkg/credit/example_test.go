package credit_test

import (
	"fmt"
	"github.com/DABronskikh/go-lesson-2_task-2/pkg/credit"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
	fmt.Println(credit.Calculate(10_000_00, 36, 20))
	// Output:
	// 3718400 33862400 133862400
	// 37184 338624 1338624
}
