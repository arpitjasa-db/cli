package mutator

import (
	"fmt"

	"github.com/databricks/cli/bundle"
	"github.com/databricks/cli/libs/dyn"
)

func (m *translatePaths) applyPipelineTranslations(b *bundle.Bundle, v dyn.Value) (dyn.Value, error) {
	fallback, err := gatherFallbackPaths(v, "pipelines")
	if err != nil {
		return dyn.InvalidValue, err
	}

	// Base pattern to match all libraries in all pipelines.
	base := dyn.NewPattern(
		dyn.Key("resources"),
		dyn.Key("pipelines"),
		dyn.AnyKey(),
		dyn.Key("libraries"),
		dyn.AnyIndex(),
	)

	for _, t := range []struct {
		pattern dyn.Pattern
		fn      rewriteFunc
	}{
		{
			base.Append(dyn.Key("notebook"), dyn.Key("path")),
			translateNotebookPath,
		},
		{
			base.Append(dyn.Key("file"), dyn.Key("path")),
			translateFilePath,
		},
	} {
		v, err = dyn.MapByPattern(v, t.pattern, func(p dyn.Path, v dyn.Value) (dyn.Value, error) {
			key := p[2].Key()
			dir, err := v.Location().Directory()
			if err != nil {
				return dyn.InvalidValue, fmt.Errorf("unable to determine directory for pipeline %s: %w", key, err)
			}

			return m.rewriteRelativeTo(b, p, v, t.fn, dir, fallback[key])
		})
		if err != nil {
			return dyn.InvalidValue, err
		}
	}

	return v, nil
}
