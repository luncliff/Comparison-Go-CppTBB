// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//
//  Note     :
//      Simple matrix abstraction of array.
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _HEAP_MATRIX_HPP_
#define _HEAP_MATRIX_HPP_

#include <memory>
#include <type_traits>

template <typename T>
class heap_matrix final
{
    static_assert(std::is_default_constructible<T>::value == true,
        "Argument type must be default constructible.");

    std::size_t row, col;
    std::unique_ptr<T[]> ptr;

public:
    // - Note
    //      Create square matrix
    explicit heap_matrix(std::size_t _width) noexcept(false)
        : row{ _width }, col{ _width }, ptr{}
    {
        ptr = std::make_unique<T[]>(_width*_width);
    }

    ~heap_matrix() noexcept = default;

    // - Note
    //      Create Row * Column size matrix
    heap_matrix(std::size_t _row, std::size_t _col) noexcept(false)
        : row{ _row }, col{ _col }, ptr{}
    {
        ptr = std::make_unique<T[]>(_row * _col)
    }

    heap_matrix(const heap_matrix &) = delete;
    heap_matrix &operator=(const heap_matrix &) = delete;

    heap_matrix(heap_matrix && rhs) noexcept :
        row{}, col{}, ptr{}
    {
        std::swap(this->ptr, rhs.ptr);
        std::swap(this->row, rhs.row);
        std::swap(this->col, rhs.col);
    }
    heap_matrix &operator=(heap_matrix && rhs) noexcept
    {
        std::swap(this->ptr, rhs.ptr);
        std::swap(this->row, rhs.row);
        std::swap(this->col, rhs.col);
        return *this;
    }

public:
    // - Note
    //      Override to access specific row
    // - Example
    //      T value = matrix[row][col]
    //      T* row  = matrix[idx]
    const T *operator[](uint32_t _row) const noexcept(false)
    {
        return ptr.get() + (_row * col);
    }
    T *operator[](uint32_t _row) noexcept(false)
    {
        return ptr.get() + (_row * col);
    }

    // - Note
    //      Length of inner array
    std::size_t size() const noexcept
    {
        return row * col;
    }
};

#endif
