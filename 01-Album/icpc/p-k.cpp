#include <iostream>

using namespace std;

int main()
{
    int r, f;
    cin >> r >> f;
    f %= 2 * r;
    if (f > r)
    {
        if ((2 * r - f) < (f - r))
            cout << "up";
        else
            cout << "down";
    }
    else
    {
        if (r - f < f)
            cout << "down";
        else
            cout << "up";
    }
}