demolist = ["kehidupan", 42, "dunis", 6, "dan", 9]
print ("demolist = ", demolist)

demolist.append ("semuanya")
print ("setelah 'semuanya' ditambahkan demolist sekarang:")
print (demolist)
print ("panjang (demolist) =", len(demolist))
print ("demolist.index(42) =", demolist.index(42))
print ("demolist [1]) = ", demolist [1])


#perulangan data dalam list

#a = 0
#while (a < (len(demolist))):
#    print ("demolist[",a,"] = ", demolist[a])
#    a = a + 1
for c in range (len(demolist)):
    print ("demolist[",c,"] =", demolist [c])

del demolist [2]
print ("setelah 'dunia' dihapus demolist sekarang:")
print (demolist)


if "kehidupan" in demolist :
    print ( "'kehidupan' ditemukan dalam demolist")
else:
    print ("'kehidupan' tidak di temukan dalam demolist")

if "amoeba" in demolist :
    print ( "'kehidupan' ditemukan dalam demolist")
if "amoeba" not in demolist:
    print ("'kehidupan' tidak di temukan dalam demolist")

list_lain = [42, 7, 0, 123]
list_lain.sort()
print ("list_lain yang sudah di urutkan", list_lain)