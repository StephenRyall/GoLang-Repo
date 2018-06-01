const factorial = x => {
  let total;
  if (x === 1) {
    return 1;
  } else {
    total = x * factorial(x - 1);
  }
  return total;
};

const additionOfFactorial = n => {
  let fac = factorial(n);
  let newFac = ("" + fac).split("").map(Number);
  console.log(newFac);
  let sum = 0;
  for (let i = 0; i < newFac.length; i++) {
    sum += newFac[i];
  }
  return sum;
};

console.log(additionOfFactorial(30));
