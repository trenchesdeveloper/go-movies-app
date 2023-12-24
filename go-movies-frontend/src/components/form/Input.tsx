import React, { forwardRef, ForwardedRef } from 'react'

interface InputProps {
    name: string
    title: string
    type: string
    placeholder?: string
    value: string
    autoComplete: string
    errorMsg?: string
    errorDiv?: string
    handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void

}

const Input = forwardRef<HTMLInputElement, InputProps>((props, ref) => {
    return (
        <div className='mb-3'>
            <label className='form-label' htmlFor={props.name}>
                {props.title}
            </label>
            <input type={props.type} className='form-control'
                id={props.name} name={props.name}
                ref={ref}
                placeholder={props.placeholder}
                onChange={props.handleChange}
                value={props.value}
                autoComplete={props.autoComplete}

            />
            <div className={props.errorDiv}>{props.errorMsg}</div>
        </div>
    )
})

export default Input