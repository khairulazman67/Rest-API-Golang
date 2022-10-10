let data = [1,2,9,9,9,9,2]

let data1 ={
    nilai : data[0],
    jumlah : 0
}
let simpanhasil = 0
for (let i=0; i<data.length; i++){
    let simpan = 0
    for(let j=0; j<data.length; j++){
        if(data[i]==data[j]){
            simpan += 1
        }
    }
    if(simpan>simpanhasil){
        simpanhasil = simpan
        data1.nilai = data[i]
        data1.jumlah = simpan
    }
}
console.log(data1.nilai)

