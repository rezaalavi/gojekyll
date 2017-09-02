package plugins

type sitemapPlugin struct{ plugin }

func init() {
	register("jekyll-sitemap", &sitemapPlugin{})
}

func (p *sitemapPlugin) PostReadSite(s Site) error {
	s.AddDocument(newTemplateDoc(s, "/sitemap.xml", sitemapTemplateSource), true)
	s.AddDocument(newTemplateDoc(s, "/robots.txt", `Sitemap: {{ "sitemap.xml" | absolute_url }}`), true)
	return nil
}

// Taken verbatim from https://github.com/jekyll/jekyll-sitemap-plugin/
const sitemapTemplateSource = `<?xml version="1.0" encoding="UTF-8"?>
{% if page.xsl %}
  <?xml-stylesheet type="text/xsl" href="{{ "/sitemap.xsl" | absolute_url }}"?>
{% endif %}
<urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  {% assign collections = site.collections | where_exp:'collection','collection.output != false' %}
  {% for collection in collections %}
    {% assign docs = collection.docs | where_exp:'doc','doc.sitemap != false' %}
    {% for doc in docs %}
      <url>
        <loc>{{ doc.url | replace:'/index.html','/' | absolute_url | xml_escape }}</loc>
        {% if doc.last_modified_at or doc.date %}
          <lastmod>{{ doc.last_modified_at | default: doc.date | date_to_xmlschema }}</lastmod>
        {% endif %}
      </url>
    {% endfor %}
  {% endfor %}

  {% assign pages = site.html_pages | where_exp:'doc','doc.sitemap != false' | where_exp:'doc','doc.url != "/404.html"' %}
  {% for page in pages %}
    <url>
      <loc>{{ page.url | replace:'/index.html','/' | absolute_url | xml_escape }}</loc>
      {% if page.last_modified_at %}
        <lastmod>{{ page.last_modified_at | date_to_xmlschema }}</lastmod>
      {% endif %}
    </url>
  {% endfor %}

  {% assign static_files = page.static_files | where_exp:'page','page.sitemap != false' | where_exp:'page','page.name != "404.html"' %}
  {% for file in static_files %}
    <url>
      <loc>{{ file.path | replace:'/index.html','/' | absolute_url | xml_escape }}</loc>
      <lastmod>{{ file.modified_time | date_to_xmlschema }}</lastmod>
    </url>
  {% endfor %}
</urlset>`
