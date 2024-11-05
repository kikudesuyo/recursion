class BankAccount {
  bankName: string;
  ownerName: string;
  savings: number;

  constructor(bankName: string, ownerName: string, savings: number) {
    this.bankName = bankName;
    this.ownerName = ownerName;
    this.savings = savings;
  }

  depositMoney(depositAmount: number) {
    if (this.savings <= 20000) {
      const commission = 100;
      this.savings += depositAmount - commission;
      return this.savings;
    }
    this.savings += depositAmount;
    return this.savings;
  }

  withdrawMoney(withdrawAmount: number) {
    if (this.savings * 0.2 < withdrawAmount) {
      this.savings -= this.savings * 0.2;
      return this.savings;
    }
    this.savings -= withdrawAmount;
    return this.savings;
  }

  pastime(days: number) {
    return this.savings + 0.25 * days;
  }
}

const user1 = new BankAccount("Chase", "Claire Simmmons", 30000);

console.log(user1.withdrawMoney(2000)); // 24000

console.log(user1.depositMoney(10000)); // 34000
console.log(user1.pastime(93)); // 36425

const user2 = new BankAccount("Bank Of America", "Remy Clay", 10000);

console.log(user2.withdrawMoney(5000)); // 8000
console.log(user2.depositMoney(12000)); // 19900
console.log(user2.pastime(505)); // 20026.25
