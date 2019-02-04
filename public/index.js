const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEwMDAwMDAwNDN9.LoExk2fFjiHABU6eKqW4af8KDDxt6RB8FQpkFhy1JAI';

d3.json(`/record?token=${token}`)
  .then(d => d.results.map(r => ({
    start: new Date(r.start),
    end: new Date(r.end)
  })))
  .then(d => draw(d))

function getHourNumber(date) {
  return Math.floor(date.getTime()/1000/60/60);
}

function draw(records) {
  const startHour = getHourNumber(new Date(`${1900 + new Date().getYear()}-01-01 00:00:00`))
  const endHour = getHourNumber(new Date(`${1901 + new Date().getYear()}-01-01 00:00:00`))
  const totalDays = Math.floor(endHour/24) - Math.floor(startHour/24);

  const svg = d3.select("body").append("svg")
    .attr('width', '100vw')
    .attr('height', 100000)
    .attr('height', `calc(100vw/24 * ${totalDays})`);
  
  const placeHolder = svg.append('g')
  _.range(totalDays).forEach(i => {
    _.range(24).forEach(j => {
      placeHolder.append('rect')
        .attr('y', d => `calc(${i} * 100vw/24 + 1)`)
        .attr('x', d => `calc(${j} * 100vw/24 + 1)`)
        .attr('height', `calc(100vw/24 - 2)`)
        .attr('width', `calc(100vw/24 - 2)`)
        .attr('fill', 'lightgrey')
    })
  })

  const nested = d3.nest()
    .key(d => getHourNumber(d.start)) // Group by hours
    .rollup(leaves => d3.sum(leaves, d => (d.end - d.start) / 1000 / 60 )) // Minutes per hour
    .entries(records);

  svg.selectAll('svg > rect')
    .data(nested)
    .enter()
    .append('rect')
    .attr('y', d => `calc(${Math.floor(d.key/24) - Math.floor(startHour/24)} * 100vw/24 + 1)`)
    .attr('x', d => `calc(${d.key % 24} * 100vw/24 + 1)`)
    .attr('height', `calc(100vw/24 - 2)`)
    .attr('width', `calc(100vw/24 - 2)`)
    .attr('fill', d => `hsl(120, 10%, ${Math.min(1 - d.value/60, 1) * 100}%)`)
}