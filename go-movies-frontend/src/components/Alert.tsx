import React from "react";

interface AlertProps {
  className?: string;
  message: string;
}

function Alert({ className, message }: AlertProps) {
  return (
    <div className={`alert ${className}`} role="alert">
      {message}
    </div>
  );
}

export default Alert;
