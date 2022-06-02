// LIB
const get = (id) => {
  return document.getElementById(id)
}

const make = (type) => {
  return document.createElement(type)
}

const insert = (parent, ...children) => {
  for (const child of children) {
    parent.append(child)
  }
}

// ON PAGE LOAD
const displayData = (data) => {
  const table = get("table")
  for (const user of data) {
    const tr = make("tr")
    const tdId = make("td")
    tdId.innerHTML = user.id
    const tdName = make("td")
    tdName = user.name
    insert(tr, tdId, tdName)
    insert(table, tr)
  }
}

const getUsers = async () => {
  const out = get("output")
  out.innerHTML = ""

  const response = await fetch("http://127.0.0.1:6789")
  if (response.ok) {
    const data = await response.json()
    displayData(data)
  } else {
    out.innerHTML = response.statusText
  }
}

getUsers()
