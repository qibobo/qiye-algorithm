package com.qiye.algorithm.sort;

public abstract class Sort {
	public void outArray(int[] a) {
		for (int i = 0; i < a.length; i++) {
			System.out.print(a[i] + ",");
		}
		System.out.print("\n");
	}

	public abstract void sort(int[] a);
}
