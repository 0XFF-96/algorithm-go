#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main()
{
    int n;
    long long ans = 0;
    cin >> n;
    vector<int> lst(n);
    for (int i = 0; i < n; i++)
    {
        cin >> lst[i];
    }
    sort(lst.begin(), lst.end());
    size_t sz = lst.size();
    int idx = sz - 2;
    for (int t = 0; t < sz / 3; t++)
    {
        // cout << lst[t] << ',';
        ans += lst[idx];
        idx -= 2;
    }
    cout << ans;
    return 0;
}

