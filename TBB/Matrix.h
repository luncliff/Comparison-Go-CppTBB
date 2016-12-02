#ifndef _MATRIX_H_
#define _MATRIX_H_

#include <memory>
#include <type_traits>

template <typename T>
class Matrix final
{
    static_assert(std::is_default_constructible<T>::value == true, 
                  "Argument type must be default constructible.");

    T* p = nullptr;
    std::size_t row{}, col{};

public:
    explicit Matrix(std::size_t _width) noexcept :
        row{ _width }, col{ _width }
    {
        p = new T[row * col]{};
    }

    Matrix(std::size_t _row, std::size_t _col) noexcept:
        row{ _row }, col{_col}
    {
        p = new T[row * col]{};
    }
    
    Matrix(Matrix&)     = delete;
    Matrix(Matrix&&)    = delete;
    Matrix& operator=(Matrix&)  = delete;
    Matrix& operator=(Matrix&&) = delete;

    ~Matrix() noexcept 
    {
        delete[] p;
    }
    
    T* operator[](int _row) 
    {
        return p + (_row * col);
    }
    
    const T* operator[](int _row) const 
    {
        return p + (_row * col);
    }

    std::size_t size() const noexcept
    {
        return row * col;
    }
};


#endif
