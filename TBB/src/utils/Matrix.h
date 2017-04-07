// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//
//  Note     :
//      Simple matrix abstraction of array.
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _MATRIX_H_
#define _MATRIX_H_

#include <memory>
#include <type_traits>

template <typename T>
class Matrix final
{
    static_assert(std::is_default_constructible<T>::value == true,
                  "Argument type must be default constructible.");

    std::unique_ptr<T> p = nullptr;
    std::size_t row{}, col{};

  public:
    // - Note
    //      Create square matrix
    explicit Matrix(std::size_t _width) noexcept : row{_width}, col{_width}
    {
        auto block = new (std::nothrow) T[row * col]{};
        p.reset(block);
    }

    // - Note
    //      Create Row * Column size matrix
    Matrix(std::size_t _row, std::size_t _col) noexcept : row{_row}, col{_col}
    {
        auto block = new (std::nothrow) T[row * col]{};
        p.reset(block);
    }

    // - Note
    //      To prevent mistake in research code, disable copy and move
    Matrix(Matrix &) = delete;
    Matrix(Matrix &&) = delete;
    Matrix &operator=(Matrix &) = delete;
    Matrix &operator=(Matrix &&) = delete;

    // - Note
    //      Override to access specific row
    // - Example
    //      T value = matrix[row][col]
    //      T* row  = matrix[idx]
    const T *operator[](int _row) const noexcept(false)
    {
        return p.get() + (_row * col);
    }
    T *operator[](int _row) noexcept(false)
    {
        return p.get() + (_row * col);
    }

    // - Note
    //      Length of inner array
    std::size_t size() const noexcept
    {
        return row * col;
    }
};

#endif
