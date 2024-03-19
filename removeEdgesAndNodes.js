function removeEdgesAndNodes(data) {
  let newObj = {}
  if (data.edges !== undefined) {
    data = removeEdges(data)
  }
}

function removeNodes(data) {

}

function removeEdges(data) {
  let edgesContent = data.edges

  let newArr = []
  edgesContent.forEach((o) => {
    newArr.push(o)
  })
  
  return newArr
}

console.log(
  removeEdgesAndNodes({
    products: {
      edges: [{ node: { title: "Havana" } }, { node: { title: "Alexandria" } }],
    },
  })
);

