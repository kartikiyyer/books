class BankAccount {
	double balance;
	public BankAccount(double balance) {
		this.balance = balance;
	}

	public void deposit(double amount) {
		balance = balance + amount;
	}

	public void withdraw(double amount) {
		if(balance >= amount) {
			balance -= amount;
		}
	}

	public void display() {
		System.out.println("Current Balance: " + balance);
	}

	public static void main(String[] args) {
		// Compulsory. Need to initialize local variables. Otherwise Compile time Error occurs.
		double balance = 0;
		if(args.length < 1) {
			System.out.println("Fewer arguments! Need to supply initial balance.");
		} else {
			try{
				balance = Double.parseDouble(args[0]);
				BankAccount ba = new BankAccount(balance);
				System.out.println("Depositing...");
				ba.deposit(10);
				ba.display();
				System.out.println("Withdrawing...");
				ba.withdraw(5);
				ba.display();
			} catch(Exception e) {
				e.printStackTrace();
			}
		}
	}
}