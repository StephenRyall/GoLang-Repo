const _ = require("lodash");

let panArray = [];
for (let s = 1; s < 99999; s++) {
  for (let b = 1; b < 99999; b++) {
    const p = s * b;
    pString = p.toString().split("");
    sString = s.toString().split("");
    bString = b.toString().split("");
    const totalLength = pString.length + sString.length + bString.length;
    if (totalLength >= 10) break;

    if (totalLength === 9) {
      const fullArray = [...pString, ...sString, ...bString];
      if (fullArray.find(str => str === "0") === "0") continue;
      if (_.uniq(fullArray).length === 9) {
        panArray.push(p);
        console.log(p, fullArray);
      }
    }
  }
}

const uniqPanArray = _.uniq(panArray);

const uniqPanArrayNum = uniqPanArray.map(str => _.toNumber(str));
console.log(_.sum(uniqPanArrayNum));
