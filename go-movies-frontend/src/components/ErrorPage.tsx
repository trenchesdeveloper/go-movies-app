import React from 'react'
import { useRouteError } from 'react-router-dom'

interface IErrorPage {
    statusText: string
    message: string

}

function ErrorPage() {
    const error: IErrorPage = useRouteError() as IErrorPage;

  return (
    <div className='container'>
        <div className="row">
          <div className="col-md-6 offset-md-3">
              <h1 className="mt-3">Oops!</h1>
              <p>Sorry, an unexpected error has occurred.</p>
              <p>
                <em>
                  {error.statusText || error.message}
                </em>
              </p>
          </div>
        </div>
    </div>
  )
}

export default ErrorPage