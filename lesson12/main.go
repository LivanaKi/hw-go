package main

import (
	"fmt"
	"sync"
)

//Конкурентно - ето разделение общих ресурсов и на их подзадачи, также роспределение их

// О чем будем говорить:
//-Группа ожидания (sync.WaitGroup);
//- Мьютексьі;
//- Гарантировано одноразовое вьіполнение
//- Race-детектор;
/*func main(){
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++{
		wg.Add(1)
		go func(id int){
			fmt.Println(id, id)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Success")
}
*/
func sum(a, b int) int{
	return a + b
}

/*
func (wg *WaitGroup) Add(delta int) - инерементирует счетчик WaitGroup на 1

func (wg *WaitGroup) Done() - декрементит счетчик на 1

func (wg *WaitGroup) Wait() - блокируется, пока счетчик WaitGroup не обнулится
*/
