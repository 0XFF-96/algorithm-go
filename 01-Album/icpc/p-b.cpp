#include <iostream>
#include <vector>
using namespace std;

int main()
{
    int n, m, k;
    cin >> n >> m >> k;
    vector<int> last(m + 1, -1);
    vector<int> costs(m + 1, 0);
    vector<vector<int>> records(m + 1);
    while (k--)
    {
        int p, c;
        cin >> p >> c;
        records[c].push_back(p);
    }
    for (int i = 1; i <= m; i++)
    {
        int sz = records[i].size();
        for (int j = 0; j < sz; j += 2)
        {
            if (j + 1 >= sz) break;
            if (records[i][j] == records[i][j + 1])
            {
                costs[i] += 100;
            }
            else
            {
                costs[i] += abs(records[i][j] - records[i][j + 1]);
            }
        }
        if (sz % 2)
        {
            costs[i] += 100;
        }
    }
    for (int i = 0; i < m; i++)
        cout << costs[i + 1] << " ";
    return 0;
}