
package rectangle

import "math"

func Diagonal(len, wid float64) float64 {
    diagonal := math.Sqrt((len * len) + (wid * wid))
    return diagonal
}
func main() {
  diagonal := diagonal(12.2,2.8)
  fmt.Printl(" diagonala este ", diagonal)


}
