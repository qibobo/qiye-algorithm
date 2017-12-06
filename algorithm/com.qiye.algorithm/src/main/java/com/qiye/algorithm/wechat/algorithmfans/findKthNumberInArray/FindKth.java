package com.qiye.algorithm.wechat.algorithmfans.findKthNumberInArray;

public class FindKth {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		FindKth f = new FindKth();
		assert (f.findKth(null, 0, 0, 0) == -1);
		assert (f.findKth(new int[] { 1 }, -1, 0, 0) == -1);
		assert (f.findKth(new int[] { 1 }, 0, -1, 0) == -1);
		assert (f.findKth(new int[] { 1 }, 2, 1, 0) == -1);
		assert (f.findKth(new int[] { 1 }, 0, 1, 4) == -1);
		assert (f.findKth(new int[] {}, 0, 1, 1) == -1);
		assert (f.findKth(new int[] { 1 }, 0, 0, 1) == 1);
		assert (f.findKth(new int[] { 2, 2, 2, 2, 2 }, 0, 4, 4) == 2);
		assert (f.findKth(new int[] { 4, 1, 2, 6, 3 }, 0, 4, 4) == 2);
		assert (f.findKth(new int[] { 4, 1, 2, 6, 3 }, 0, 4, 1) == 6);
		assert (f.findKth(new int[] { 4, 1, 2, 6, 3 }, 0, 4, 2) == 4);
		assert (f.findKth(new int[] { 4, 1, 2, 6, 3 }, 0, 4, 3) == 3);
		assert (f.findKth(new int[] { 4, 1, 2, 6, 3 }, 0, 4, 4) == 2);
		assert (f.findKth(new int[] { 4, 1, 2, 6, 3 }, 0, 4, 5) == 1);
		assert (f.findKth(new int[] { 4, 1, 2, 6, 3 }, 0, 4, 6) == -1);
	}

	public int findKth(int[] a, int s, int e, int k) {
		if (a == null || a.length == 0 || s < 0 || e < 0 || s > e || k > e - s + 1) {
			return -1;
		}
		for (int i = (e - 1) / 2; i >= s; i--) {
			this.heapAdjust(a, i, e);
		}
		int j;
		for (j = 1; j <= k; j++) {
			int tmp = a[e - j + 1];
			a[e - j + 1] = a[s];
			a[s] = tmp;
			this.heapAdjust(a, s, e - j);
		}
		return a[e - (j - 1) + 1];
	}

	private void heapAdjust(int[] a, int s, int e) {
		int tmp = a[s];
		int parent = s;
		for (int i = 2 * s + 1; i <= e; i = i * 2 + 1) {
			if (i + 1 <= e && a[i + 1] > a[i]) {
				i = i + 1;
			}
			if (tmp >= a[i]) {
				break;
			}
			a[parent] = a[i];
			parent = i;
		}
		a[parent] = tmp;
	}

}
