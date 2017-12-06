package com.qiye.algorithm.sort.heap;

import com.qiye.algorithm.sort.Sort;

public class ArrayHeapSort  extends Sort{

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int[] a = { 33, 22, 1, 4, 2, 123, 56, 888 };
		Sort sort = new ArrayHeapSort();
		sort.sort(a);
		sort.outArray(a);
	}

	public void sort(int[] a) {
		for (int i = (a.length-1) / 2; i >= 0; i--) {
			this.adjust(a, i, a.length - 1);
		}
		outArray(a);
		for (int j = a.length - 1; j > 0; j--) {
			int tmp = a[0];
			a[0] = a[j];
			a[j] = tmp;
			this.adjust(a, 0, j - 1);
			outArray(a);
		}
	}

	private void adjust(int[] a, int start, int end) {
		int tmp = a[start];
		for (int i = 2 * start + 1; i <= end; i = 2 * i + 1) {
			if (i + 1 <= end && a[i + 1] > a[i]) {
				i++;
			}
			if (tmp >= a[i]) {
				break;
			}
			a[start] = a[i];
			start = i;
		}
		a[start] = tmp;
	}

}
