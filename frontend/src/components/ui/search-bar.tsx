"use client";

import { useState, useCallback, useEffect, useRef } from "react";
import { Search, X, Loader2 } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";

interface SearchBarProps {
  value?: string;
  onChange?: (value: string) => void;
  onSearch?: (value: string) => void;
  placeholder?: string;
  className?: string;
  autoFocus?: boolean;
  showSuggestions?: boolean;
  suggestions?: string[];
  onSuggestionClick?: (suggestion: string) => void;
  isLoading?: boolean;
  size?: "default" | "lg";
}

export function SearchBar({
  value = "",
  onChange,
  onSearch,
  placeholder = "Cari dataset, artikel...",
  className,
  autoFocus = false,
  showSuggestions = false,
  suggestions = [],
  onSuggestionClick,
  isLoading = false,
  size = "default",
}: SearchBarProps) {
  const [inputValue, setInputValue] = useState(value);
  const [showDropdown, setShowDropdown] = useState(false);
  const wrapperRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    setInputValue(value);
  }, [value]);

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (wrapperRef.current && !wrapperRef.current.contains(event.target as Node)) {
        setShowDropdown(false);
      }
    }
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleSubmit = useCallback(
    (e: React.FormEvent) => {
      e.preventDefault();
      onSearch?.(inputValue);
      setShowDropdown(false);
    },
    [inputValue, onSearch]
  );

  const handleChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const val = e.target.value;
      setInputValue(val);
      onChange?.(val);
      if (showSuggestions && val.length >= 2) {
        setShowDropdown(true);
      } else {
        setShowDropdown(false);
      }
    },
    [onChange, showSuggestions]
  );

  const handleClear = useCallback(() => {
    setInputValue("");
    onChange?.("");
    onSearch?.("");
  }, [onChange, onSearch]);

  const handleSuggestionClick = useCallback(
    (suggestion: string) => {
      setInputValue(suggestion);
      onSuggestionClick?.(suggestion);
      setShowDropdown(false);
      onSearch?.(suggestion);
    },
    [onSuggestionClick, onSearch]
  );

  const lgSize = size === "lg";

  return (
    <div ref={wrapperRef} className={cn("relative", className)}>
      <form onSubmit={handleSubmit} className="relative">
        <Search
          className={cn(
            "absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground",
            lgSize ? "h-5 w-5" : "h-4 w-4"
          )}
        />
        <Input
          value={inputValue}
          onChange={handleChange}
          placeholder={placeholder}
          autoFocus={autoFocus}
          className={cn(
            "pl-10 pr-10",
            lgSize && "h-12 text-base"
          )}
        />
        {isLoading ? (
          <Loader2 className="absolute right-3 top-1/2 -translate-y-1/2 h-4 w-4 animate-spin text-muted-foreground" />
        ) : inputValue ? (
          <button
            type="button"
            onClick={handleClear}
            className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground transition-colors"
          >
            <X className="h-4 w-4" />
          </button>
        ) : null}
      </form>

      {/* Suggestions dropdown */}
      {showSuggestions && showDropdown && suggestions.length > 0 && (
        <div className="absolute z-50 mt-1 w-full rounded-md border bg-popover shadow-md">
          <ul className="py-1">
            {suggestions.map((suggestion, index) => (
              <li key={index}>
                <button
                  type="button"
                  onClick={() => handleSuggestionClick(suggestion)}
                  className="w-full px-4 py-2 text-left text-sm hover:bg-accent transition-colors flex items-center gap-2"
                >
                  <Search className="h-3.5 w-3.5 text-muted-foreground shrink-0" />
                  {suggestion}
                </button>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

// Hero search variant
export function HeroSearchBar(props: Omit<SearchBarProps, "size">) {
  return (
    <SearchBar
      {...props}
      size="lg"
      className="max-w-2xl mx-auto"
    />
  );
}
