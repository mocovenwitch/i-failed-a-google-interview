package searchandsort.firstbadversion;

public class FirstBadVersion {

    // RESULT: 15.55%, 70.85%
    public int firstBadVersion0(int n) {
        return binarySearch(1, n);
    }

    private int binarySearch(int min, int max) {
        int mid = (max - min) / 2 + min;
        boolean currentVersion = isBadVersion(mid);
        boolean prevVersion = isBadVersion(mid - 1);

        if (!prevVersion && currentVersion) {
            // found first one, (false, true)
            return mid;
        } else if (prevVersion && currentVersion) {
            // (true, true) means too big
            return binarySearch(min, mid - 1);
        } else if (!prevVersion && !currentVersion) {
            // (false, false) mean too small
            return binarySearch(mid + 1, max);
        } else {
            throw new IllegalArgumentException("Bad version inputs!");
        }
    }

    // Optimised:
    public int firstBadVersion(int n) {
        if (n == 0) return 0;
        int min = 1, max = n;

        while (min <= max) {
            int mid = min + (max - min) / 2;

            if (isBadVersion(mid)) {
                max = mid - 1;
            } else {
                min = mid + 1;
            }
        }

        return min;
    }

    boolean isBadVersion(int version) {
        return version >= 87;
    }
}
