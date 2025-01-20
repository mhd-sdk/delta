import { Star, StarFilled } from '@carbon/icons-react';
import { Button, ContainedList, ContainedListItem, InlineLoading, Search } from '@carbon/react';
import { css } from '@emotion/css';
import { debounce } from 'lodash';
import { useEffect, useMemo, useRef, useState } from 'react';
import { GetAssets } from '../../../wailsjs/go/main/App';
import { alpaca, persistence } from '../../../wailsjs/go/models';
import { useAppData } from '../../hooks/useAppData';

interface Props {
  onSelect: (ticker: string) => void;
  onClose: () => void;
}

export const TickerList = ({ onSelect, onClose }: Props): JSX.Element => {
  const [assets, setAssets] = useState<alpaca.Asset[]>([]);
  const [loading, setLoading] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState<alpaca.Asset[]>([]);

  const { appData, onSave } = useAppData();
  const favoriteTickers = appData.favoriteTickers;

  const containerRef = useRef<HTMLDivElement>(null);
  const searchInputRef = useRef<HTMLInputElement>(null);

  const handleChange = useMemo(
    () =>
      debounce((event) => {
        setSearchTerm(event.target.value);
        setLoading(false);
      }, 500),
    []
  );

  useEffect(() => {
    const tickersFound = assets.filter((listItem) => listItem.symbol.toLowerCase().includes(searchTerm.toLowerCase()));
    const namesFound = assets.filter((listItem) => listItem.name.toLowerCase().includes(searchTerm.toLowerCase()));
    if (tickersFound.length > 0) {
      setSearchResults(tickersFound);
    } else {
      setSearchResults(namesFound);
    }
  }, [assets, searchTerm]);

  useEffect(() => {
    const fetchAssets = async () => {
      const res = await GetAssets();
      setAssets(res);
      setSearchResults(res);
    };
    fetchAssets();
  }, []);

  useEffect(() => {
    if (searchInputRef.current) {
      searchInputRef.current.focus();
    }
  }, []);

  const isFavorite = (ticker: string) => favoriteTickers.includes(ticker);

  const Fav = (ticker: string) => (
    <Button
      kind="ghost"
      renderIcon={isFavorite(ticker) ? StarFilled : Star}
      iconDescription={isFavorite(ticker) ? 'Remove from favorites' : 'Add to favorites'}
      tooltipAlignment="end"
      hasIconOnly
      onClick={() => handleAddToFav(ticker)}
    />
  );

  const renderList = () => {
    if (loading) {
      return (
        <ContainedListItem className={styles.listItem}>
          <InlineLoading status="active" description="Loading tickers" />
        </ContainedListItem>
      );
    }
    if (searchTerm.length === 0) {
      return favoriteTickers.map((ticker, key) => (
        <ContainedListItem action={Fav(ticker)} className={styles.listItem} key={key} onClick={() => onSelect(ticker)}>
          {ticker}
        </ContainedListItem>
      ));
    }
    if (searchResults.length === 0) {
      return (
        <ContainedListItem className={styles.listItem}>
          <InlineLoading status="error" description="No results found" />
        </ContainedListItem>
      );
    }
    if (searchTerm.length !== 0) {
      return searchResults.slice(0, 50).map((ticker, key) => (
        <ContainedListItem action={Fav(ticker.symbol)} className={styles.listItem} key={key} onClick={() => onSelect(ticker.symbol)}>
          {ticker.symbol}
        </ContainedListItem>
      ));
    }
  };

  const handleAddToFav = (ticker: string) => {
    if (isFavorite(ticker)) {
      onSave({ ...appData, favoriteTickers: favoriteTickers.filter((fav) => fav !== ticker) } as persistence.AppData);
    } else {
      onSave({ ...appData, favoriteTickers: [...favoriteTickers, ticker] } as persistence.AppData);
    }
  };

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (containerRef.current && !containerRef.current.contains(event.target as Node)) {
        onClose();
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [onClose]);

  return (
    <div ref={containerRef} className={styles.layout}>
      <div className={styles.body}>
        <ContainedList
          kind="disclosed"
          label={
            <Search
              placeholder=""
              labelText=""
              defaultValue={searchTerm}
              onChange={(e) => {
                setLoading(true);
                handleChange(e);
              }}
              className={styles.search}
              closeButtonLabelText="Clear search input"
              size="sm"
              ref={searchInputRef} // Attach ref to the search input
            />
          }
          size="sm"
        >
          {renderList()}
        </ContainedList>
      </div>
    </div>
  );
};

const styles = {
  box: css`
    padding: 5px;
    height: 50px;
  `,
  search: css`
    input {
      text-transform: uppercase !important;
    }
  `,
  listItem: css`
    cursor: pointer;
  `,
  body: css`
    flex: 1;
    overflow-y: auto;
    overflow-x: hidden;
  `,
  layout: css`
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 170px;
    overflow-x: hidden;
    max-height: 300px;
    .cds--contained-list__header {
      padding: 0;
    }
  `,
  actions: css`
    margin-left: auto;
  `,
  moreResults: css`
    margin-top: 1rem;
    font-size: 0.875rem;
    color: gray;
  `,
};
