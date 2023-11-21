# Weighted Strings

Weighted Strings merupakan suatu algoritma untuk mencari bobot huruf dari suatu string yang diberikan. 

Sampel:
Input:
Diberikan sebuah string 'abbcccd' dengan queries [1, 3, 9, 8]. Queries digunakan untuk menentukan status dari angka di dalam queries dengan aturan:
1. Apabila angka di queries bernilai sama dengan bobot karakter/substring maka return Yes.
2. Apabila angka di queries bernilai beda dengan bobot karakter/substring maka return No.
Pembobotan substring dan karakter:
a = 1
b = 2
bb = 4
c = 3
cc = 6
ccc = 9
d = 4 
Output: ['Yes', 'Yes', 'Yes', 'No']
Penjelasan: Dari pembobotan substring dan karakter yang dimiliki 'abbcccd' maka status dari queries yang ditentukan yaitu ['Yes', 'Yes', 'Yes', 'No']. 

kompleksitas dari kode tersebut adalah O(n * n^2). dikarenakan pada fungsi calculateWeight untuk menghitung seberapa bobot dari character yang di masukkan dan pada fungsi weightedStrings untuk pengecekan dari weight tersebut sudah sesuai dari inputQueries yang diharapkan dengan metoda backtracking.