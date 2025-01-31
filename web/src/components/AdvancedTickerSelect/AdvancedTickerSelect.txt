// import {
//   DataTable,
//   Pagination,
//   Table,
//   TableBody,
//   TableCell,
//   TableHead,
//   TableHeader,
//   TableRow,
//   TableToolbar,
//   TableToolbarContent,
//   TableToolbarSearch,
// } from '@carbon/react';
// import { css } from '@emotion/css';
// import { debounce } from 'lodash';
// import { useEffect, useMemo, useState } from 'react';
// import { alpaca } from '../../../wailsjs/go/models';

// interface Props {
//   onBack: () => void;
//   disabled: boolean;
//   onSelect: (symbol: string) => void;
// }

// export const TickerAdvancedSelector = ({ onBack, disabled, onSelect }: Props): JSX.Element => {
//   const [tickers, setTickers] = useState<alpaca.Asset[]>([]);
//   const [filteredSymbols, setFilteredSymbols] = useState<alpaca.Asset[]>([]);
//   const [currentPage, setCurrentPage] = useState(1);
//   const [pageSize, setPageSize] = useState(50);
//   const [searchTerm, setSearchTerm] = useState('');
//   const [isLoading, setIsLoading] = useState(true);

//   // Débounce la recherche pour réduire les recalculs fréquents
//   const debouncedSetSearchTerm = useMemo(
//     () =>
//       debounce((value: string) => {
//         setSearchTerm(value.toLowerCase());
//       }, 300),
//     []
//   );

//   const headers = [
//     { key: 'name', header: 'Name' },
//     { key: 'symbol', header: 'Symbol' },
//     { key: 'type', header: 'Type' },
//   ];

//   useEffect(() => {
//     const fetchSymbols = async () => {
//       setIsLoading(true);
//       const res = await GetSymbols();
//       setTickers(res);
//       setFilteredSymbols(res);
//       setIsLoading(false);
//     };

//     fetchSymbols();
//   }, []);

//   // Filtrage des symboles par recherche
//   useEffect(() => {
//     const results = tickers.filter((symbol) => symbol.name.toLowerCase().includes(searchTerm));
//     setFilteredSymbols(results);
//     setCurrentPage(1); // Réinitialiser à la première page lors d'une nouvelle recherche
//   }, [searchTerm, tickers]);

//   const displayedSymbols = filteredSymbols.slice((currentPage - 1) * pageSize, currentPage * pageSize);

//   return (
//     <div className={styles.layout}>
//       <TableToolbar>
//         <TableToolbarContent>
//           <TableToolbarSearch onChange={(_, value) => value !== undefined && debouncedSetSearchTerm(value)} placeholder="Search symbols..." />
//         </TableToolbarContent>
//       </TableToolbar>

//       {isLoading ? (
//         <div
//           className={css`
//             text-align: center;
//             padding: 2rem;
//           `}
//         >
//           Loading...
//         </div>
//       ) : (
//         <DataTable
//           rows={displayedSymbols.map((symbol) => ({
//             id: symbol.symbol,
//             name: symbol.name,
//             symbol: symbol.symbol,
//           }))}
//           headers={headers}
//         >
//           {({ rows, headers, getHeaderProps, getRowProps }) => (
//             <Table className={styles.widthHeight100}>
//               <TableHead className={styles.tableHead}>
//                 <TableRow>
//                   {headers.map((header) => (
//                     <TableHeader {...getHeaderProps({ header })} key={header.key}>
//                       {header.header}
//                     </TableHeader>
//                   ))}
//                 </TableRow>
//               </TableHead>
//               <TableBody>
//                 {rows.map((row) => (
//                   <TableRow {...getRowProps({ row })} onClick={() => onSelect(row.id)} key={row.id}>
//                     {row.cells.map((cell) => (
//                       <TableCell key={cell.id}>{cell.value}</TableCell>
//                     ))}
//                   </TableRow>
//                 ))}
//               </TableBody>
//             </Table>
//           )}
//         </DataTable>
//       )}

//       <Pagination
//         page={currentPage}
//         pageSize={pageSize}
//         pageSizes={[10, 20, 50, 100]}
//         totalItems={filteredSymbols.length}
//         onChange={({ page, pageSize }) => {
//           setCurrentPage(page);
//           setPageSize(pageSize);
//         }}
//       />
//     </div>
//   );
// };

// const styles = {
//   widthHeight100: css`
//     width: 100%;
//     height: 100%;
//   `,
//   tableHead: css`
//     position: sticky !important;
//   `,
//   layout: css`
//     display: flex;
//     flex-direction: column;
//     height: 100%;
//   `,
//   actions: css`
//     display: flex;
//     justify-content: flex-end;
//     padding-top: 1rem;
//   `,
// };
