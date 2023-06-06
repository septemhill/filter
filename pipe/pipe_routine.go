package pipe

type mapFunc[A, B any] func(A) (B, error)

func iterFlatMap[A, B any](reader <-chan A, writer chan<- B, fn mapFunc[A, B]) {
	for elem := range reader {
		data, err := fn(elem)
		if err != nil {
			continue
		}
		writer <- data
	}
	close(writer)
}

func pushElements[A any](list []A, writer chan<- A) {
	for _, elem := range list {
		writer <- elem
	}
	close(writer)
}

func PipeRoutine2[A, B, C any](
	list []A,
	f1 func(A) (B, error),
	f2 func(B) (C, error)) []C {
	p1 := make(chan A, 5)
	p2 := make(chan B, 5)
	out := make(chan C, 5)

	go pushElements(list, p1)
	go iterFlatMap(p1, p2, f1)
	go iterFlatMap(p2, out, f2)

	result := make([]C, 0)
	for elem := range out {
		result = append(result, elem)
	}

	return result
}

func PipeRoutine3[A, B, C, D any](
	list []A,
	f1 func(A) (B, error),
	f2 func(B) (C, error),
	f3 func(C) (D, error)) []D {
	p1 := make(chan A, 5)
	p2 := make(chan B, 5)
	p3 := make(chan C, 5)
	out := make(chan D, 5)

	go pushElements(list, p1)
	go iterFlatMap(p1, p2, f1)
	go iterFlatMap(p2, p3, f2)
	go iterFlatMap(p3, out, f3)

	result := make([]D, 0)
	for elem := range out {
		result = append(result, elem)
	}

	return result
}

func PipeRoutine4[A, B, C, D, E any](
	list []A,
	f1 func(A) (B, error),
	f2 func(B) (C, error),
	f3 func(C) (D, error),
	f4 func(D) (E, error)) []E {
	p1 := make(chan A, 5)
	p2 := make(chan B, 5)
	p3 := make(chan C, 5)
	p4 := make(chan D, 5)
	out := make(chan E, 5)

	go pushElements(list, p1)
	go iterFlatMap(p1, p2, f1)
	go iterFlatMap(p2, p3, f2)
	go iterFlatMap(p3, p4, f3)
	go iterFlatMap(p4, out, f4)

	result := make([]E, 0)
	for elem := range out {
		result = append(result, elem)
	}

	return result
}

func PipeRoutine5[A, B, C, D, E, F any](
	list []A,
	f1 func(A) (B, error),
	f2 func(B) (C, error),
	f3 func(C) (D, error),
	f4 func(D) (E, error),
	f5 func(E) (F, error)) []F {
	p1 := make(chan A, 5)
	p2 := make(chan B, 5)
	p3 := make(chan C, 5)
	p4 := make(chan D, 5)
	p5 := make(chan E, 5)
	out := make(chan F, 5)

	go pushElements(list, p1)
	go iterFlatMap(p1, p2, f1)
	go iterFlatMap(p2, p3, f2)
	go iterFlatMap(p3, p4, f3)
	go iterFlatMap(p4, p5, f4)
	go iterFlatMap(p5, out, f5)

	result := make([]F, 0)
	for elem := range out {
		result = append(result, elem)
	}

	return result
}
