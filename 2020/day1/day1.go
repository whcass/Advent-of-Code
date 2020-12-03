package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	input := "1587\n1407\n1717\n1596\n1566\n1752\n1925\n1847\n1716\n1726\n1611\n1628\n1853\n1864\n1831\n1942\n1634\n1964\n1603\n1676\n1256\n1906\n1655\n1790\n1666\n1470\n1540\n1544\n1100\n1447\n1384\n1464\n1651\n1572\n907\n1653\n1265\n1510\n1639\n1818\n376\n1378\n1132\n1750\n1491\n1788\n1882\n1779\n1640\n1586\n1525\n1458\n1994\n1782\n1412\n1033\n1416\n1813\n1520\n1968\n715\n1396\n1745\n1506\n1024\n1798\n1870\n1615\n1957\n1718\n1349\n1983\n1387\n1738\n1588\n1321\n1160\n1907\n1861\n1940\n1475\n2004\n1852\n1760\n1608\n1028\n1820\n1495\n1811\n1737\n1417\n1316\n1087\n1803\n1595\n1346\n1971\n1692\n1678\n1330\n1480\n1097\n1898\n1973\n1567\n1733\n1336\n1381\n1327\n1670\n1436\n1989\n1334\n89\n1862\n1715\n1743\n1967\n1765\n1402\n1729\n1749\n1671\n1196\n1650\n1089\n1814\n1783\n1225\n1823\n1746\n2009\n1886\n1748\n1481\n1739\n1912\n1663\n1668\n1314\n1594\n705\n1449\n1731\n1487\n1648\n1466\n1317\n1979\n1799\n1926\n1703\n1656\n1978\n2005\n1865\n1982\n1951\n1892\n1713\n1744\n1598\n1606\n1583\n1895\n1804\n1430\n1816\n1364\n1575\n1918\n1431\n1812\n1471\n1797\n928\n1934\n1156\n94\n1563\n1909\n1453\n1392\n1427\n1819\n1524\n1695\n1866\n2008\n1413\n1698\n1051\n1707\n1904\n1681\n1541\n1621\n1421\n1809\n1576"
	sliceData := strings.Split(input, "\n")
	part1(sliceData)
	part2(sliceData)
	fmt.Println("[*] Done.")
}

func part2(sliceData []string) {
	for i := 0; i < len(sliceData); i++ {
		for j := 0; j < len(sliceData); j++ {
			for k := 0; k < len(sliceData); k++ {
				a, _ := strconv.Atoi(sliceData[i])
				b, _ := strconv.Atoi(sliceData[j])
				c, _ := strconv.Atoi(sliceData[k])
				d := a + b + c

				if d == 2020 {
					fmt.Printf("[*] Part 2: %d\n", a*b*c)
					return
				}
			}
		}
	}
}

func part1(sliceData []string) {
	for i := 0; i < len(sliceData); i++ {
		for j := i; j < len(sliceData); j++ {
			a, _ := strconv.Atoi(sliceData[i])
			b, _ := strconv.Atoi(sliceData[j])
			c := a + b
			if c == 2020 {
				fmt.Printf("[*] Part 1: %d\n", a*b)
				return
			}
		}
	}
}
