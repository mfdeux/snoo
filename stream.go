package snoo

import (
	"sync"
	"time"
)

type CommentStream struct {
	client   *Client
	Comments chan *Comment
	done     chan struct{}
	group    *sync.WaitGroup
	onlyNew  bool
	linkID   string
}

func (s *CommentStream) Stop() {
	close(s.done)
	// block until the retry goroutine stops
	s.group.Wait()
}

func (s *CommentStream) Start() *CommentStream {
	var seenComments []*Comment
	go func() {
		comments, err := s.client.GetLinkComments(s.linkID)
		if err != nil {
			s.Stop()
			return
		}
		for _, comment := range comments {
			seenComments = append(seenComments, comment)
			if !s.onlyNew {
				s.Comments <- comment
			}
		}
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case _ = <-ticker.C:
				comments, err := s.client.GetLinkComments(s.linkID)
				if err != nil {
					s.Stop()
					return
				}
				for _, comment := range comments {
					isSeen := false
					for _, seenComment := range seenComments {
						if comment.ID == seenComment.ID {
							isSeen = true
						}
					}
					if !isSeen {
						seenComments = append(seenComments, comment)
						s.Comments <- comment
					}
				}
			case <-s.done:
				return
			}
		}
	}()
	return s
}

type LinkStream struct {
	client    *Client
	Links     chan Link
	done      chan struct{}
	group     *sync.WaitGroup
	onlyNew   bool
	subreddit string
}

func (s *LinkStream) Stop() {
	close(s.done)
	// Scanner does not have a Stop() or take a done channel, so for low volume
	// streams Scan() blocks until the next keep-alive. Close the resp.Body to
	// escape and stop the stream in a timely fashion.

	// block until the retry goroutine stops
	s.group.Wait()
}

func (s *LinkStream) Start() *LinkStream {
	var seenLinks []Link
	go func() {
		links, err := s.client.GetNewLinks(s.subreddit)
		if err != nil {
			s.Stop()
			return
		}
		for _, link := range links {
			seenLinks = append(seenLinks, link)
			if !s.onlyNew {
				s.Links <- link
			}
		}
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case _ = <-ticker.C:
				links, err := s.client.GetNewLinks(s.subreddit)
				if err != nil {
					s.Stop()
					return
				}
				for _, link := range links {
					isSeen := false
					for _, seenLink := range seenLinks {
						if link.ID == seenLink.ID {
							isSeen = true
						}
					}
					if !isSeen {
						seenLinks = append(seenLinks, link)
						s.Links <- link
					}
				}
			case <-s.done:
				return
			}
		}
	}()
	return s
}

// StreamLinks is a streaming device
func (c *Client) StreamLinks(subreddit string, onlyNew bool) *LinkStream {
	s := &LinkStream{
		client:    c,
		Links:     make(chan Link),
		done:      make(chan struct{}),
		group:     &sync.WaitGroup{},
		onlyNew:   onlyNew,
		subreddit: subreddit,
	}

	return s
}

// StreamLinkComments is a streaming device
func (c *Client) StreamLinkComments(linkID string, onlyNew bool) *CommentStream {
	s := &CommentStream{
		client:   c,
		Comments: make(chan *Comment),
		done:     make(chan struct{}),
		group:    &sync.WaitGroup{},
		onlyNew:  onlyNew,
		linkID:   linkID,
	}

	return s
}
