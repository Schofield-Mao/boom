package heap


func swap(a []string, l int, r int){
	temp := a[l]
	a[l] = a[r]
	a[r] = temp
}

func shiftDown(a []string, size int, i int, m map[string]int){
	for {
		i=2*i+1
		if i>=size{
			break
		}
		if m[a[i]]>=m[a[i+1]]{
			i = i+1
		}
		if m[a[i]]<=m[a[(i-1)/2]]{
			swap(a,i,(i-1)/2)
		}else{
			break
		}
	}
}

func shiftUp(a []string, i int,  m map[string]int){
	for {
		if i<=0{
			break
		}
		if m[a[i]]<=m[a[(i-1)/2]]{
			swap(a,i,(i-1)/2)
		}
		i = (i-1)/2
	}
}

func insert(v string, a *[]string, m map[string]int){
	*a = append(*a,v)
	size := len(*a)
	shiftUp(*a,size-1,m)
}

func remove(a *[]string, m map[string]int){
	index := 0
	(*a)[0] = (*a)[len(*a)-1]
	shiftDown(*a,len(*a)-1,index,m)
	*a = append((*a)[0:len(*a)-1])
}

// func makeHeap(a []string,m map[string]int){
// 	for i:=(len(a)-1)/2;i>=0;i--{
// 		shiftDown(a,len(a)-1,i,m)
// 	}
// }

// func heapSort(a []string,m map[string]int){
// 	MakeHeap(a,m)
// 	for size := len(a)-1;size>0;size--{
// 		swap(a,0,size)
// 		shiftDown(a,size-1,0,m)
// 	} 
// }

func GetTopK(m map[string]int, num_k int) []string{
  topK := make([]string,0)
  for k,v := range m{
    if len(topK)<num_k{
	  insert(k,&topK,m)
	}else if m[topK[0]] < v{
	  remove(&topK,m)
      insert(k,&topK,m)
	}
  }
	return topK 
}