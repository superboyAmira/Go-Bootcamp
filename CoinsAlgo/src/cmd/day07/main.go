package main

import (
	mincoins "day07/internal/minCoins"
	"log"
	"os"
	"runtime/pprof"
)

// Main func to get pprof report
func main() {
	f, err := os.Create("cpu.prof")
    if err != nil {
        log.Fatal("could not create CPU profile: ", err)
    }
    defer f.Close() 

    if err := pprof.StartCPUProfile(f); err != nil {
        log.Fatal("could not start CPU profile: ", err)
    }
    defer pprof.StopCPUProfile()

	tests := []struct {
		input []int
		sum   int

		expected []int
	}{
		{[]int{}, 15, []int{}},
		{[]int{1, 5, 10}, 13, []int{10, 1, 1, 1}},
		{[]int{1, 1, 1, 5, 5, 10}, 13, []int{10, 1, 1, 1}},
		{[]int{1, 1, 1, 5, 5, 12}, 13, []int{12, 1}},
		{[]int{9, 1, 2, 3, 4, 5}, 17, []int{9, 5, 3}},
		{mincoins.GenerateSlice(10000, 1), 10000000, mincoins.GenerateSlice(10000000, 1)},
		{[]int{
			1, 7, 13, 21, 28, 35, 43, 50, 56, 64,
			71, 79, 83, 97, 101, 109, 117, 123, 130, 137,
			142, 149, 155, 160, 168, 175, 183, 190, 197, 202,
			209, 215, 223, 230, 239, 245, 251, 260, 267, 273,
			281, 289, 297, 302, 310, 318, 326, 332, 340, 349,
			355, 362, 370, 378, 384, 391, 398, 406, 412, 419,
			428, 433, 441, 449, 456, 463, 471, 478, 485, 492,
			499, 507, 513, 521, 528, 536, 543, 551, 559, 565,
			573, 581, 588, 596, 602, 610, 618, 625, 631, 639,
			647, 654, 661, 668, 676, 684, 692, 698, 706, 713,
			721, 729, 735, 742, 750, 759, 766, 773, 779, 787,
			794, 803, 811, 818, 826, 833, 841, 847, 855, 862,
			870, 877, 884, 892, 900, 907, 915, 921, 929, 936,
			945, 953, 960, 967, 975, 983, 990, 998, 1003, 1010,
			1022, 1031, 1038, 1045, 1053, 1060, 1067, 1075, 1081, 1089,
			1097, 1103, 1112, 1119, 1127, 1133, 1141, 1148, 1155, 1162,
			1171, 1178, 1185, 1193, 1200, 1207, 1215, 1222, 1228, 1237,
			1243, 1251, 1258, 1266, 1272, 1280, 1287, 1295, 1302, 1310,
			1318, 1325, 1333, 1340, 1347, 1354, 1362, 1369, 1377, 1383,
			1391, 1399, 1406, 1414, 1420, 1428, 1437, 1444, 1451, 1458,
			1466, 1472, 1481, 1487, 1495, 1502, 1510, 1517, 1524, 1531,
			1540, 1546, 1553, 1562, 1570, 1577, 1584, 1590, 1599, 1606,
			1613, 1621, 1628, 1635, 1642, 1650, 1657, 1665, 1672, 1678,
			1686, 1692, 1701, 1708, 1716, 1723, 1730, 1737, 1744, 1752,
			1760, 1767, 1773, 1780, 1787, 1795, 1803, 1810, 1817, 1825,
			1832, 1840, 1846, 1853, 1861, 1868, 1875, 1883, 1890, 1897,
			1905, 1912, 1920, 1926, 1934, 1941, 1949, 1956, 1963, 1970,
			1978, 1985, 1993, 2000,
		}, 1500050, []int{}},
	}

	for _, t := range tests {
		mincoins.MinCoins2(t.sum, t.input)
	}
}
