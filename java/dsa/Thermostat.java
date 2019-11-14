class Thermostat {
	int currentTemp;
	int desiredTemp;

	Thermostat(int currentTemp) {
		this.currentTemp = currentTemp;
	}

	void setDesiredTemp(int desiredTemp) {
		this.desiredTemp = desiredTemp;
	}

	boolean furnaceOn() {
		if(currentTemp < desiredTemp) {
			System.out.println("Furnace is ON now!");
			try {
				Thread.sleep(1000);
			} catch (Exception e) {
				e.printStackTrace();
			}
			currentTemp++;
			return true;
		}
		return false;
	}

	void furnaceOff() {
		if(currentTemp == desiredTemp) {
			System.out.println("Furnace is OFF now!");
		}
	}

	public static void main(String[] args) {
		System.out.println("Hello World!");
		Thermostat t = new Thermostat(50);
		t.setDesiredTemp(55);
		while(t.furnaceOn()) {

		}
		t.furnaceOff();
	}
}