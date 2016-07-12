package others

import "fmt"

func Demo1()  {
	z:=37
	pi:=&z
	ppi:=&pi
	fmt.Println(z,*pi,**ppi)
	**ppi++
	fmt.Println(z,*pi,**ppi)
	fmt.Println(ppi)
}

func Demo2()  {
	i:=9
	j:=5
	i,j,product:=swapAndProduct2(i,j)
	fmt.Println(i,j,product);
}

func Demo3()  {
	s:=[]string{"A","B","C","D","E","F","G"}
	t:=s[:5]
	u:=s[3:len(s)-1]
	fmt.Println(s,t,u)
	u[1]="x"
	fmt.Println(s,t,u)
}

func Demo4()  {
	amounts:=[]float64{237.81,261.87,273.93,279.99,281.07,303.17,231.47,227.33,209.23,197.09}
	sum:=0.0
	for _,amount:=range amounts{
		sum+=amount
	}
	fmt.Println(sum)
	for i:=range amounts{
		amounts[i]*=1.05
	}
	fmt.Println(amounts)
}

type Product struct{
	name string
	price float64
}

func (product Product)String()string  {
	return fmt.Sprintf("%s (%.2f)",product.name,product.price)
}

func Demo5()  {
	products:=[]*Product{{"Spanner",3.99},{"Wrench",2.49},{"Screwdriver",1.99}}
	fmt.Println(products)
	for _,product:=range products{
		product.price+=0.50
	}
	fmt.Println(products);
}
func swapAndProduct2(x,y int)(int,int,int)  {
	if x>y{
		x,y=y,x
	}
	return x,y,x*y
}