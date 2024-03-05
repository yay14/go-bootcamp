package main

import "fmt"

type Matrix struct {
	no_of_rows    int
	no_of_columns int
	arr           [][]int
}

func (matrix *Matrix) get_no_of_rows() int {
	return matrix.no_of_rows
}

func (matrix *Matrix) get_no_of_columns() int {
	return matrix.no_of_columns
}

func (matrix *Matrix) set_val_at_ind(row, col, new_val int) {
	if row < matrix.get_no_of_rows() &&
		col < matrix.get_no_of_columns() {
		matrix.arr[row][col] = new_val
	} else {
		fmt.Printf("Index out of bound\n")
	}
}

func (matrix1 Matrix) add_matrix(matrix2 Matrix) {

	if matrix1.get_no_of_columns() == matrix2.no_of_rows &&
		matrix1.get_no_of_rows() == matrix2.no_of_columns {

		temp_two_d := make([][]int, matrix1.get_no_of_rows())

		for i := range matrix1.get_no_of_rows() {

			temp_one_d := make([]int, matrix1.get_no_of_columns())

			for j := range matrix1.get_no_of_columns() {
				temp_one_d[j] = matrix1.arr[i][j] + matrix2.arr[i][j]
			}

			temp_two_d[i] = temp_one_d
		}

		matrix_sum := Matrix{
			matrix1.get_no_of_rows(),
			matrix1.get_no_of_columns(),
			temp_two_d,
		}

		matrix_sum.print_matrix_in_json_format()

	} else {
		fmt.Printf("Dimention mismatch")
	}
}

func (matrix *Matrix) print_matrix_in_json_format() {
	fmt.Printf("[\n")

	for i := range matrix.get_no_of_rows() {

		fmt.Printf("  [")

		for j := range matrix.get_no_of_columns() {

			fmt.Printf("%d", matrix.arr[i][j])

			if j+1 != matrix.get_no_of_columns() {
				fmt.Printf(", ")
			}

		}
		fmt.Printf("],\n")
	}

	fmt.Printf("]")
}

func main() {
	total_rows := 3
	total_col := 3

	arr1 := make([][]int, total_rows)
	arr2 := make([][]int, total_rows)

	for i := range total_rows {
		temp_arr1 := make([]int, total_col)
		temp_arr2 := make([]int, total_col)

		for j := range total_col {
			temp_arr1[j] = j + 1
			temp_arr2[j] = temp_arr1[j] * 2
		}

		arr1[i] = temp_arr1
		arr2[i] = temp_arr2
	}

	matrix1 := Matrix{
		total_rows,
		total_col,
		arr1,
	}

	matrix2 := Matrix{
		total_rows,
		total_col,
		arr2,
	}

	fmt.Printf("No of rows in matrix 1 : %d\n", matrix1.get_no_of_rows())
	fmt.Printf("No of rows in matrix 1 : %d\n", matrix1.get_no_of_columns())
	matrix1.set_val_at_ind(0, 1, 23)
	fmt.Printf("Sum of matrix 1 & 2\n")
	matrix1.add_matrix(matrix2)

}
