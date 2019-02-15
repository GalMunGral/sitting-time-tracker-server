// TEST TOKEN
const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEwMDAwMDAwMDJ9.mFDbR_B_Ehm7c6Tqreb3W6q-yLU6xWD_YKpe2wk8ZdY'

d3.json(`/record?token=${token}`)
  .then(d => d.results.map(r => ({
    start: new Date(r.start.replace(/-/g, '/')), // WebKit does not support 'yyyy-MM-dd'.
    end: new Date(r.end.replace(/-/g, '/')) // Must convert to 'yyyy/MM/dd'!!!
  })))
  .then(d => draw(d))

function getHourNumber(date) {
  return Math.floor(date.getTime()/1000/60/60);
}

function draw(records) {
  const startHour = getHourNumber(new Date(`${1900 + new Date().getYear()}/01/01 00:00:00`))
  const endHour = getHourNumber(new Date(`${1901 + new Date().getYear()}/01/01 00:00:00`))
  const totalDays = Math.floor(endHour/24) - Math.floor(startHour/24);

  let vw = window.innerWidth;

  const svg = d3.select("body").append("svg")
    // attr('width', '100vw')
    .attr('width', vw)
    // .attr('height', `calc(100vw/24 * ${totalDays})`);
    .attr('height', vw / 24 * totalDays);
  
  const placeHolder = svg.append('g')
  _.range(totalDays).forEach(i => {
    _.range(24).forEach(j => {
      placeHolder.append('rect')
        // .attr('y', d => `calc(${i} * 100vw/24 + 1)`)
        // .attr('x', d => `calc(${j} * 100vw/24 + 1)`)
        // .attr('height', `calc(100vw/24 - 2)`)
        // .attr('width', `calc(100vw/24 - 2)`)
        // .attr('fill', 'lightgrey')
        .attr('y', i * vw / 24 + 1)
        .attr('x', j * vw / 24 + 1)
        .attr('height', vw / 24 - 2)
        .attr('width', vw / 24 - 2)
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
    // .attr('y', d => `calc(${Math.floor(d.key/24) - Math.floor(startHour/24)} * 100vw/24 + 1)`)
    // .attr('x', d => `calc(${d.key % 24} * 100vw/24 + 1)`)
    // .attr('height', `calc(100vw/24 - 2)`)
    // .attr('width', `calc(100vw/24 - 2)`)
    .attr('y', d => (Math.floor(d.key/24) - Math.floor(startHour/24)) * vw / 24 + 1)
    .attr('x', d => ((d.key - startHour) % 24) * vw / 24 + 1)
    .attr('height', vw / 24 - 2)
    .attr('width', vw / 24 - 2)
    // .attr('fill', d => `hsl(120, 10%, ${Math.min(1 - d.value/60, 1) * 100}%)`)
    .attr('fill', d => {
      return hslToRgb(0.333, 0.1, Math.max(1 - d.value/60, 0))
      // return 'black'
    })
}

/**
 * Code copied from https://gist.github.com/mjackson/5311256
 * h, s, l are within range [0, 1]
 */
function hslToRgb(h, s, l) {
  var r, g, b;

  if (s == 0) {
    r = g = b = l; // achromatic
  } else {
    function hue2rgb(p, q, t) {
      if (t < 0) t += 1;
      if (t > 1) t -= 1;
      if (t < 1/6) return p + (q - p) * 6 * t;
      if (t < 1/2) return q;
      if (t < 2/3) return p + (q - p) * (2/3 - t) * 6;
      return p;
    }

    var q = l < 0.5 ? l * (1 + s) : l + s - l * s;
    var p = 2 * l - q;

    r = hue2rgb(p, q, h + 1/3);
    g = hue2rgb(p, q, h);
    b = hue2rgb(p, q, h - 1/3);
  }
  // WKWebKit only recognizes rgb(int, int, int) for svg fill attribute!!!
  return `rgb(${Math.floor(r * 255)}, ${Math.floor(g * 255)}, ${Math.floor(b*255)}`
}