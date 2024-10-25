package main


import "os"


func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func ft_putnbr(n int) {
	if n < 0 {
		ft_putchar('-')
		n = -n
	}
	if n >= 10 {
		ft_putnbr(n / 10)
	}
	ft_putchar(byte(n%10 + '0')) 
}

func	ft_puthex(numhexa uint16, control int){
	if numhexa >= 16{
		ft_puthex(numhexa / 16, control)
	}
	if control == 1{
		ft_putchar("0123456789abcdef"[numhexa % 16])
	}else{
		ft_putchar("0123456789ABCEDF"[numhexa % 16])
	}
	
}
func ft_printf(format string, args ...interface{}) {

	argIndex := 0

	for i := 0; i < len(format); i++ {
		if format[i] == '%' {
			i++
			if i < len(format) {
				switch format[i] {
				case 's', 'c':
					if argIndex < len(args) {
						if str, ok := args[argIndex].(string); ok {
							for j := 0; j < len(str); j++ {
								ft_putchar(str[j])
							}
						}
						argIndex++
					}
				case 'd', 'i':
					if argIndex < len(args) {
						if num, ok := args[argIndex].(int); ok {
							ft_putnbr(num)
						}
						argIndex++
					}
				case 'x':
					if argIndex < len(args){
						if numhexa, ok := args[argIndex].(uint16); ok{
							ft_puthex(numhexa, 1);
						}
						argIndex++
					}
				case 'X':
					if argIndex < len(args){
						if numhexa, ok := args[argIndex].(uint16); ok{
							ft_puthex(numhexa, 0);
						}
						argIndex++
					}
				default:
					ft_putchar(format[i])
				}
			}
		} else {
			ft_putchar(format[i])
		}
	}
}

func main() {
	nome := "Daniel"
	ft_printf("%s\n", nome) 
	ft_printf("O número é: %d\n meu nome é: %s\n hexa: %x\n HEXA %X\n", 42,  "João", uint16(10), uint16(10))
}
