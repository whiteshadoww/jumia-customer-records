import {useEffect, useState} from "react";

export default function Main() {

    const [customers, updateCustomers] = useState([])
    const [country, setCountry] = useState('')
    const [state, setState] = useState('')
    const [pageSize, setPageSize] = useState(10)
    const [page, setPage] = useState(1)

    let url = `http://127.0.0.1:3001/api/v1/customers?country=${country}&page_size=${pageSize}&page=${page}`
    if (state !== '')
        url += `&state=${state}`

    useEffect(() => {
        fetch(url)
            .then(response => response.json())
            .then(r => {
                if (r.status === 200 && r.code === "SUCCESS") {
                    updateCustomers(r.data ?? [])
                }
            })
    }, [country, state, pageSize, page])


    return (
        <>
            <div className="min-h-full">

                <header className="bg-white shadow">
                    <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
                        <h1 className="text-3xl font-bold text-gray-900">Customer Phones</h1>
                    </div>
                </header>
                <main>
                    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
                        {/* Replace with your content */}
                        <div className="px-4 py-6 sm:px-0">

                            <div className="relative overflow-x-auto shadow-md sm:rounded-lg">

                                <div className="grid xl:grid-cols-2 xl:gap-6">
                                    <div className="relative z-0 mb-6 w-full group">
                                        <div className="p-4 col-6">
                                            <label htmlFor="table-search" className="sr-only">Search</label>
                                            <div className="relative mt-1">
                                                <div
                                                    className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
                                                    <svg className="w-5 h-5 text-gray-500 dark:text-gray-400"
                                                         fill="currentColor" viewBox="0 0 20 20"
                                                         xmlns="http://www.w3.org/2000/svg">
                                                        <path fill-rule="evenodd"
                                                              d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                                                              clip-rule="evenodd"></path>
                                                    </svg>
                                                </div>
                                                <input type="text" id="table-search"
                                                       onChange={e => setCountry(e.target.value)}
                                                       className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-80 pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                                       placeholder="Search for country"/>


                                            </div>
                                        </div>
                                    </div>
                                    <div className="relative z-0 mb-6 w-full group pr-2">
                                        <label htmlFor="countries"
                                               className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400">Select
                                            State</label>
                                        <select id="countries"
                                                onChange={(ee) => setState(ee.target.value)}
                                                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
                                            <option value={''}>All</option>
                                            <option value={'false'}>In valid</option>
                                            <option value={'true'}>valid</option>
                                        </select>
                                    </div>
                                </div>

                                <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                                    <thead
                                        className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                                    <tr>

                                        <th scope="col" className="px-6 py-3">
                                            Customer Name
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Phone
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Country
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            Country Code
                                        </th>
                                        <th scope="col" className="px-6 py-3">
                                            State
                                        </th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    {customers.map((m) => <tr
                                        className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">

                                        <th scope="row"
                                            className="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                                            {m.name}
                                        </th>
                                        <td className="px-6 py-4">
                                            {m.phone}
                                        </td>
                                        <td className="px-6 py-4">
                                            {m.country}
                                        </td>
                                        <td className="px-6 py-4">
                                            {m.country_code}
                                        </td>
                                        <td className={
                                            'px-6 py-4'
                                        }>
                                            {m.state ? "Valid" : "Not Valid"}
                                        </td>
                                    </tr>)
                                    } </tbody>
                                </table>
                            </div>


                            <div className="flex flex-col items-center">

                                <div className="inline-flex mt-2 xs:mt-0">
                                    <button
                                        disabled={page === 1}
                                        onClick={() => setPage(state => --state)}
                                        className="py-2 px-4 text-sm font-medium text-white bg-gray-800 rounded-l hover:bg-gray-900 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                                        Prev
                                    </button>
                                    <button
                                        onClick={() => setPage(state => ++state)}
                                        className="py-2 px-4 text-sm font-medium text-white bg-gray-800 rounded-r border-0 border-l border-gray-700 hover:bg-gray-900 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                                        Next
                                    </button>
                                </div>
                            </div>


                        </div>
                        {/* /End replace */}
                    </div>
                </main>
            </div>
        </>
    )
}
