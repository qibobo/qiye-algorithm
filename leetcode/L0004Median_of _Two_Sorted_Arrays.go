package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len12 := len(nums1) + len(nums2)
	if (len12)%2 == 0 {
		return (float64(FindKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, len12/2)) + float64(FindKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, len12/2+1))) / 2
	} else {
		return float64(FindKth(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, len12/2+1))
	}
}
func FindKth(a []int, as, ae int, b []int, bs, be int, k int) int {
	lena := ae - as + 1
	lenb := be - bs + 1
	// must be the first check
	if lena > lenb {
		return FindKth(b, bs, be, a, as, ae, k)
	}
	if lena == 0 {
		return b[bs+k-1]
	}
	if k == 1 {
		if a[as] > b[bs] {
			return b[bs]
		} else {
			return a[as]
		}
	}

	var indexa, indexb int
	if lena < k/2 {
		indexa = ae
		indexb = bs + k - lena - 1
	} else {
		indexa = as + k/2 - 1
		indexb = bs + k - k/2 - 1
	}
	if a[indexa] < b[indexb] {
		return FindKth(a, indexa+1, ae, b, bs, be, k-(indexa-as+1))
	} else {
		return FindKth(a, as, ae, b, indexb+1, be, k-(indexb-bs+1))
	}
}
