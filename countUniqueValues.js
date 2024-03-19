function countUniqueValues(array) {
   let values = new Map()
   let uniqueCount = 0

   array.forEach((e) => {
      if (values.get(e) === undefined) {
        uniqueCount++
        values.set(e,1)
      } 
   })
   return uniqueCount
}


console.log(countUniqueValues([-2, 0, 1, 1, 3, 3, 4])); // 5
console.log(countUniqueValues([1,1,1,1,1,2,-1,2])); // 5
console.log(countUniqueValues([])); // 0
